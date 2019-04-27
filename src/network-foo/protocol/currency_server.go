package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	curr "github.com/vladimirvivien/go-networking/currency/lib0"
)

var currencies = curr.Load("../github.com/vladimirvivien/go-networking/currency/data.csv")

// Implements a simple lookup service over TCP.
// Program loads ISO currency info using package lib0
// and use a simple text-based protocol to interact witt
// the client and send the data.
//
// Clients send search requests as a textual command:
//
// GET <currency, country, or code>
//
// When the server receives, it parses the request
// and used the request to search the list of currencies.
// The search result is then printed line-by-line back to client.
//
// This server implements a streaming strategy when receiving data from client
// to avoid dropping data when the request is larger than the internal buffer.
//
// Lerverage the capability of the net.Conn object implementing io.Reader
// which allows us to stream data.
//

func main() {
	var addr string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr]")
	flag.Parse()

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to create listener: %v \n", err)
	}

	defer listen.Close()
	log.Println("**** TCP based Currency Server ***")
	log.Printf("Service started: %s\n", addr)

	// Standard server loop
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			if err := conn.Close(); err != nil {
				log.Println("Failed to close connection: ", err)
			}
			continue
		}
		log.Println("Connected to ", conn.RemoteAddr())

		// Hand off incoming request to handleConnection routine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Error closing connection: ", err)
		}
	}()

	if _, err := conn.Write([]byte("Connected... \nUsage: GET <currency, country, or code>\n")); err != nil {
		log.Println("Error writting: ", err)
		return
	}

	// loop to stay connected with client
	for {
		// buffer for the client command
		var cmdLine []byte

		// stream data using 4-byte chunks until io.EOF
		// chunks are kept small to demo streaming feature
		for {
			chunk := make([]byte, 4)
			n, err := conn.Read(chunk)
			if err != nil {
				if err == io.EOF {
					cmdLine, _ = appendBytes(cmdLine, chunk[:n])
					break
				}
				log.Println("connection read error: ", err)
				return
			}

			if cmdLine, err = appendBytes(cmdLine, chunk[:n]); err == io.EOF {
				break
			}
		}

		cmd, param := parseCommand(string(cmdLine))
		switch strings.ToUpper(cmd) {
		case "GET":
			results := curr.Find(currencies, param)
			if len(results) == 0 {
				if _, err := conn.Write([]byte("Nothing found\n")); err != nil {
					log.Println("failed to write: ", err)
				}
				continue
			}

			// send each currency info as a line to client
			for _, cur := range results {
				s := fmt.Sprintf("%s %s %s %s\n", cur.Name, cur.Code, cur.Number, cur.Country)
				_, err := conn.Write([]byte(s))
				if err != nil {
					log.Println("failed to write response: ", err)
					return
				}
			}
		default:
			if _, err := conn.Write([]byte("Invalid command \n")); err != nil {
				log.Println("failed to write: ", err)
				return
			}
		}
	}
}

func parseCommand(cmdLine string) (cmd, param string) {
	parts := strings.Split(cmdLine, " ")
	if len(parts) != 2 {
		return "", ""
	}
	cmd = strings.TrimSpace(parts[0])
	param = strings.TrimSpace(parts[1])
	return cmd, param
}

// for each byte b in src
// 	if b == '\n' return io.EOF
//  else append(dest, b)
func appendBytes(dest, src []byte) ([]byte, error) {
	for _, b := range src {
		if b == '\n' {
			return dest, io.EOF
		}
		dest = append(dest, b)
	}
	return dest, nil
}

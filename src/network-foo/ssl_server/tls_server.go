package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

var currencies = curr.Load("../../github.com/vladimirvivien/go-networking/currency/data.csv")

// Implements a simple lookup service over TCP.
// Program loads ISO currency info using package lib
// and use a simple JSON protocol to interact with
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
// This server uses the enconding/json package Encoder/Decoder types
// 	which are accept an io.Writer and io.Reader
//

func main() {
	var addr string
	var certFile string
	var keyFile string
	flag.StringVar(&addr, "e", ":4343", "service endpoint [ip addr]")
	flag.StringVar(&certFile, "cert", "../ssl/db-win-cert.pem", "tls server certificate")
	flag.StringVar(&keyFile, "key", "../ssl/db-win-key.pem", "tls server key")
	flag.Parse()

	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	listen, err := tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		log.Println("Failed to create listener: ", err)
		os.Exit(1)
	}

	defer listen.Close()
	log.Println("**** TCP based Currency Server ***")
	log.Printf("Service started: %s\n", addr)

	// Delay to sleep when accepting fails with a temp error
	acceptDelay := 10 * time.Microsecond
	acceptCount := 0

	// Standard server loop
	for {
		conn, err := listen.Accept()
		if err != nil {
			switch e := err.(type) {
			case net.Error:
				if e.Temporary() {
					if acceptCount > 5 {
						log.Printf("Unable to connect after %d retries: %v", acceptCount, err)
						return
					}
					acceptDelay *= 2
					acceptCount++
					time.Sleep(acceptDelay)
					continue
				}
			default:
				fmt.Println(err)
				conn.Close()
				continue
			}
			acceptDelay = 10 * time.Microsecond
			acceptCount = 0
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

	// its initial request or loose the connection.
	if err := conn.SetDeadline(time.Now().Add(time.Second * 45)); err != nil {
		log.Println("failed to set deadline:", err)
		return
	}
	// command-loop
	for {
		// Decode incoming data into curr.CurrencyRequest
		decoder := json.NewDecoder(conn)
		var req curr.CurrencyRequest

		if err := decoder.Decode(&req); err != nil {
			switch e := err.(type) {
			case net.Error:
				if e.Timeout() {
					fmt.Println("deadline reached, disconnecting...")
				}
				log.Println("Network error:", e)
				return
			default:
				if e == io.EOF {
					log.Println("Closing connection:", e)
					return
				}
				encoder := json.NewEncoder(conn)
				if encErr := encoder.Encode(&curr.CurrencyError{Error: e.Error()}); encErr != nil {
					fmt.Println("Failed error encoding:", encErr)
					return
				}
				continue
			}
		}

		results := curr.Find(currencies, req.Get)
		encoder := json.NewEncoder(conn)
		if err := encoder.Encode(&results); err != nil {
			switch e := err.(type) {
			case net.Error:
				fmt.Println("failed to send response:", e)
				return
			default:
				if encErr := encoder.Encode(&curr.CurrencyError{Error: e.Error()}); encErr != nil {
					fmt.Println("Failed to send error:", encErr)
					return
				}
				continue
			}
		}

		if err := conn.SetDeadline(time.Now().Add(time.Second * 45)); err != nil {
			fmt.Println("failed to set deadline:", err)
			return
		}
	}
}

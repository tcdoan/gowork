package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

const prompt = "currency"

func main() {

	var addr string
	flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
	flag.Parse()

	// create a dialer to configure its settings instead
	// of using the default dialer from net.Dial() function.
	dialer := &net.Dialer{
		Timeout:   time.Second * 300,
		KeepAlive: time.Minute * 5,
	}

	var (
		conn           net.Conn
		err            error
		connTries      = 0
		connMaxRetries = 3
		connSleepRetry = time.Second * 1
	)

	for connTries < connMaxRetries {
		fmt.Println("creating connection socket to", addr)
		conn, err = dialer.Dial("tcp", addr)

		if err != nil {
			fmt.Printf("failed to dial %s. Error: %v\n", addr, err)
			switch nerr := err.(type) {
			case net.Error:
				if nerr.Temporary() {
					connTries++
					fmt.Println("trying again in:", connSleepRetry)
					time.Sleep(connSleepRetry)
					continue
				}
				fmt.Println("unable to recover")
				os.Exit(1)
			default:
				os.Exit(1)
			}
		}
		break
	}

	// did we get a connection

	if conn == nil {
		fmt.Println("failed to create a connection to ", addr)
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connected to currency service: ", addr)
	var param string

	for {
		fmt.Print(prompt, "> ")
		_, err = fmt.Scanf("%s", &param)

		if err != nil {
			fmt.Println("Usage: <search string or *>")
			continue
		}

		req := curr.CurrencyRequest{Get: param}

		// Send request:
		// use json encoder to encode value of type curr.CurrencyRequest
		// and stream it to the server via net.Conn.
		if err := json.NewEncoder(conn).Encode(&req); err != nil {
			switch err := err.(type) {
			case net.Error:
				fmt.Println("failed to send request:", err)
				os.Exit(1)
			default:
				fmt.Println("failed to encode request:", err)
				continue
			}
		}

		// Display response
		var currencies []curr.Currency
		err = json.NewDecoder(conn).Decode(&currencies)
		if err != nil {
			switch err := err.(type) {
			case net.Error:
				fmt.Println("failed to receive response:", err)
				os.Exit(1)
			default:
				fmt.Println("failed to decode response:", err)
				continue
			}
		}

		fmt.Println(currencies)
	}
}

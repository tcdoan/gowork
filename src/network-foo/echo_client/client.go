package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost:4040", "echo service endpoint")
	flag.Parse()
	text := flag.Arg(0)

	raddr, _ := net.ResolveTCPAddr("tcp", host)
	conn, _ := net.DialTCP("tcp", nil, raddr)
	conn.Write([]byte(text))
	b := make([]byte, 256)
	n, _ := conn.Read(b)
	fmt.Printf("%s \n", string(b[:n]))
}

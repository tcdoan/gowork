package main

import (
	"flag"
	"net"
)

func main() {
	var host string
	flag.StringVar(&host, "host", ":8080", "echo service endpoint")
	flag.Parse()
	endpoint, _ := net.ResolveTCPAddr("tcp", host)
	listen, _ := net.ListenTCP("tcp", endpoint)
	defer listen.Close()
	for {
		conn, _ := listen.AcceptTCP()
		go procoess(conn)
	}
}

func procoess(conn *net.TCPConn) {
	defer conn.Close()
	b := make([]byte, 256)
	n, _ := conn.Read(b)
	conn.Write(b[:n])
}

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "us.pool.ntp.org:123", "NTP Host")
	flag.Parse()

	// serverAddr is UDPAddr struct represents remote end point
	remoteAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		log.Fatalf("Failed to UDP resolve %s. Error: %v\n", host, err)
	}

	// 48-byte request to NTP server. First byte 0x1B
	req := make([]byte, 48)
	req[0] = 0x1B

	// Holding incoming datagram with time values from
	// us.pool.ntp.org:123
	resp := make([]byte, 48)

	// Setup net.UDPConn with net.DialUDP()
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatalf("Failed to connect %s. Error: %v\n", host, err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed closing connection: %v \n", err)
			os.Exit(1)
		}
	}()

	fmt.Printf("Time from %s \n", conn.RemoteAddr())

	// send request
	if _, err = conn.Write(req); err != nil {
		fmt.Printf("Failed to send request: %v \n", err)
		os.Exit(1)
	}

	// blocking
	read, err := conn.Read(resp)
	if err != nil {
		log.Fatalf("Failed to receive server response: %v \n", err)
	}

	if read != 48 {
		log.Fatalf("Failed to real all 48 bytes from server: %v \n", err)
	}

	// NTP resp is big-endian []byte (LSB [0...47] MSB)
	// with a 64-bit value containing the server time in seconds
	// First 32-bits are seconds. Last 32-bit are fractional.
	// Extracting number of seconds from [0...[40:43]...47]
	// The number of secs since NTP epoch (1900)
	secs := binary.BigEndian.Uint32(resp[40:])
	frac := binary.BigEndian.Uint32(resp[44:])
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	now := float64(secs) - offset
	fmt.Printf("%v\n", time.Unix(int64(now), int64(frac)))
}

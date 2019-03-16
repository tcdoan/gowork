package main

import (
	"io"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])
	io.Copy(os.Stdout, f)
}

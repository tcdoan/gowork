package main

import (
	"flag"
	"fmt"
	"tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "-temp 20C")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

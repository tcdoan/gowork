package main

import (
	"display"
)

func main() {
	var x interface{} = 1
	display.Display("x", x)
	display.Display("x", &x)
}

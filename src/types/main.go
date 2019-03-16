package main

import (
	"io"
	"os"
)

func main() {
	var f io.Writer = os.Stdout

	if f2, ok := f.(*os.File); ok {
		f2.WriteString("Hello \n")
	}

}

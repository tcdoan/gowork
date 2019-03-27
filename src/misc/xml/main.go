package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	dec := xml.NewDecoder(resp.Body)
	var stack []string

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v \n", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[2:]) {
				fmt.Printf("%s: %s \n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}

		x = x[1:]
	}
	return false
}

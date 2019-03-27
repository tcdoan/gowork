package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Printf("%s \n", <-ch)
	}
	fmt.Printf("Elapsed time: %.2f seconds.", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, _ := http.Get(url)
	bytes, _ := io.Copy(ioutil.Discard, resp.Body)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Elapsed time: %.2f, %7d, %s", secs, bytes, url)
}

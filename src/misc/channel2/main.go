package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")

	c := fanIn(joe, ann)

	for i := 0; ; i++ {
		fmt.Println(<-c)
	}

}

func fanIn(c1 <-chan string, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func(m string, c chan string) {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- fmt.Sprintf("%s %d", m, i)
		}
	}(msg, c)
	return c
}

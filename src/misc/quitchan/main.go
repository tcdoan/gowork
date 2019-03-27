package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str string
}

func main() {
	joeQuit := make(chan string)
	annQuit := make(chan string)

	joe := boring("Joe", joeQuit)
	ann := boring("Ann", annQuit)

	c := fanIn(joe, ann)

	for i := 0; ; i++ {
		m1 := <-c
		fmt.Println(m1.str)

		m2 := <-c
		fmt.Println(m2.str)

		if i == 10 {
			joeQuit <- "bye Joe"
			annQuit <- "bye Ann"
			break
		}
	}
	fmt.Printf("Joe say: %q \n", <-joeQuit)
	fmt.Printf("Ann say: %q \n", <-annQuit)
}

func fanIn(c1 <-chan message, c2 <-chan message) <-chan message {
	c := make(chan message)
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

func boring(msg string, quit chan string) <-chan message {
	c := make(chan message)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- message{str: fmt.Sprintf("%s %d", msg, i)}:
			case <-quit:
				time.Sleep(2 * time.Second)
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

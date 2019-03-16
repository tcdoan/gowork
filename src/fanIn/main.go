package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str  string
	wait chan bool
}

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	c := fanIn(joe, ann)

	for i := 0; i < 5; i++ {
		m1 := <-c
		fmt.Println(m1.str)

		m2 := <-c
		fmt.Println(m2.str)
		m2.wait <- true
		m1.wait <- true
	}
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

func boring(msg string) <-chan message {
	c := make(chan message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			m := message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			c <- m
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

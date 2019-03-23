package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const MyLuckyNumber = 6
	luck := false
	numBadLucks := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for !luck {
		x := r.Intn(10)
		if x == MyLuckyNumber {
			luck = true
		} else {
			numBadLucks++
		}
	}

	fmt.Printf("Number of bad lucks %d", numBadLucks)
}

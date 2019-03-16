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

	for !luck {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := r.Intn(10)

		fmt.Println(x, x == MyLuckyNumber)
		if x == MyLuckyNumber {
			luck = true
		} else {
			numBadLucks++
		}
	}

	fmt.Printf("Number of bad lucks %d", numBadLucks)
}

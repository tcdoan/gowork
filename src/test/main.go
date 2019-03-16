package main

import "fmt"

func main() {
	var errors = [...]string{
		5:  "no such process", // ESRCH
		10: "no such process", // ESRCH
	}

	fmt.Println(len(errors))
}

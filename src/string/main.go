package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "aà"
	r := []rune(s)

	for i, r := range s {
		fmt.Printf("%d \t %q \t %d \t %x \n", i, r, r, r)
	}

	fmt.Println(unsafe.Sizeof(r[0]) + unsafe.Sizeof(r[1]))
	fmt.Println(len(s))
}

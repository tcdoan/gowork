package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	x1 bool
	x2 int16
	x3 float64
}

type s2 struct {
	x1 float64
	x2 int16
	x3 bool
}

func main() {
	x1 := s1{}
	fmt.Println(unsafe.Alignof(x1.x1))
	fmt.Println(unsafe.Alignof(x1.x2))
	fmt.Println(unsafe.Alignof(x1.x3))

	// 0
	fmt.Println(unsafe.Offsetof(x1.x1))
	// 2
	fmt.Println(unsafe.Offsetof(x1.x2))
	// 8
	fmt.Println(unsafe.Offsetof(x1.x3))

	fmt.Printf("%d \n", float64bits(1.0))
}

func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

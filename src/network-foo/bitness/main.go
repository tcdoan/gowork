package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	var x uint16
	//x = F0FF
	x = 61695
	fmt.Printf("y: %.16b \n", x)

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, &x)
	if err != nil {
		fmt.Println(err)
	}

	buf2 := new(bytes.Buffer)
	err = binary.Write(buf2, binary.BigEndian, &x)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("buf: %b \n", buf)
	fmt.Printf("buf2: %b \n", buf2)
}

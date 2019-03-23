package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! \n", user.Name)
	fmt.Printf("Type in monkey command \n")
	repl.Start(os.Stdin, os.Stdout)
}

package main

func main() {
	c := make(chan string)
	c <- string([]byte("Hi there!"))
}

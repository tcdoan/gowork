package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type chineseBot struct{}

func main() {
	eb := englishBot{}
	cb := chineseBot{}
	printGreeting(eb)
	printGreeting(cb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (chineseBot) getGreeting() string {
	return "嗨，您好"
}

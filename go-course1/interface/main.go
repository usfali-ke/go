package main

import "fmt"

type bot interface {
	getGreetings() string
}

type engBot struct{}
type espBot struct{}

func main() {

	en := engBot{}
	esp := espBot{}

	printGreetings(en)
	printGreetings(esp)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (e engBot) getGreetings() string {
	return "Hi there!"
}

func (s espBot) getGreetings() string {
	return "Hola!"
}

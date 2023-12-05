package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Exiting the program.")

}

// send
func foo(c chan<- int) {
	c <- 89
}

// receive
func bar(c <-chan int) {
	fmt.Println("Value received from the channel is", <-c)
}

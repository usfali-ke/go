package main

import "fmt"

func main() {
	c1 := make(chan int)   //bi-directional chan 0 send/receive
	c2 := make(chan<- int) // send to chan
	c3 := make(<-chan int) // receive from chan

	// c1 <- 39 // value sent to chan c1
	// fmt.Println(<-c1) // receive from chan c1 to stdio
	// c2 <- 43
	// fmt.Println(<-c3)

	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c3)
}

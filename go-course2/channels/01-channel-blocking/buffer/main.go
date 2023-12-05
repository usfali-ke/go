package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 43
	c <- 45

	fmt.Println(<-c)
	fmt.Println(<-c)
}

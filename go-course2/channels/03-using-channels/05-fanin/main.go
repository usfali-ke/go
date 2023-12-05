package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	fanin := make(chan int)

	go send(c1, c2)

	go receive(c1, c2, fanin)

	for i := range fanin {
		fmt.Println(i)
	}
}

// send

func send(c1, c2 chan<- int) {
	for i := 1; i < 15; i++ {
		if i%2 == 0 {
			c1 <- i
		} else {
			c2 <- i
		}
	}
	close(c1)
	close(c2)
}

// receive
func receive(c1, c2 <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := range c1 {
			fanin <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range c2 {
			fanin <- i
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

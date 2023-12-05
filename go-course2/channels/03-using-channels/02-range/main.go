package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 50; i < 89; i++ {
			c <- i
		}
		close(c)
	}()

	for j := range c {
		fmt.Println(j)
	}
}

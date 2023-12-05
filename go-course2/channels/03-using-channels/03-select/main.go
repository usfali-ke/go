package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send2()

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Value from even channel\t", v)
		case v := <-odd:
			fmt.Println("Value from odd channel\t", v)
		case <-quit:
			fmt.Println("quiting..")
			return
		}
	}
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 15; i <= 50; i++ {
		if i%2 == 0 {
			// fmt.Println(i)
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- true
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	even := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan bool)

// 	go send(even, odd, quit)

// 	receive(even, odd, quit)

// 	fmt.Println("about to exit")
// }

// // send channel
// func send(even, odd chan<- int, quit chan<- bool) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			even <- i
// 		} else {
// 			odd <- i
// 		}
// 	}
// 	quit <- true
// }

// // receive channel
// func receive(even, odd <-chan int, quit <-chan bool) {
// 	for {
// 		select {
// 		case v := <-even:
// 			fmt.Println("the value received from the even channel:", v)
// 		case v := <-odd:
// 			fmt.Println("the value received from the odd channel:", v)
// 		case <-quit:
// 			return
// 		}
// 	}
// }

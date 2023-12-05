package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx := context.Background()
	fmt.Println("Context:\t", ctx)
	fmt.Println("Context error 1:\t", ctx.Err())
	fmt.Println("Goroutines 1:\t", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)

			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("Context error 2:\t", ctx.Err())
	fmt.Println("Goroutines 2:\t", runtime.NumGoroutine())

	fmt.Println("about to cancel...")
	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("Context error 3:\t", ctx.Err())
	fmt.Println("Goroutines 3:\t", runtime.NumGoroutine())

}

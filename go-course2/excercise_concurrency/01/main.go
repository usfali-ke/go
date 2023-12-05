package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("starting cpu", runtime.NumCPU())
	fmt.Println("starting routines", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hello there from the fist routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello there from the second routine")
		wg.Done()
	}()
	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid routines", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Exiting the program..")

	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end routines", runtime.NumGoroutine())
}

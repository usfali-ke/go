package main

import "fmt"

func main() {
	f := increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// closure
// A Go closure is an anonymous nested function which retains bindings to variables defined outside the body of the closure.
func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

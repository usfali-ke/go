package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", substract)
	fmt.Printf("%T\n", doMath)

	x := doMath(20, 5, add)
	fmt.Println(x)

	y := doMath(46, 12, substract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}

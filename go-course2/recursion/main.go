package main

import "fmt"

func main() {

	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}

func factorial(n int) int {
	fmt.Println("this is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// recursion above can also be achieved with a loop
func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--

	}
	return total
}

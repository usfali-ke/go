package main

import "fmt"

func main() {
	fmt.Println("4 + 5 =", addNum(4, 5))
	fmt.Println("10 + 2 =", addNum(10, 2))
	fmt.Println("12 + 15 =", addNum(12, 15))
}

func addNum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
		// sum += v

	}
	return sum
}

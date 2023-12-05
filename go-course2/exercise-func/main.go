package main

import "fmt"

func main() {
	num1 := foo(56)
	fmt.Println(num1)

	num2, val := bar(34, "testval")
	fmt.Println(num2, val)

	intVal := []int{23, 27, 10, 20, 15, 5}
	// unfurl int slice intval
	sum := foo1(intVal...)
	fmt.Println(sum)

	sum = foo1(12, 66, 9, 1, 2)
	fmt.Println(sum)

	barSUm := bar1(intVal)
	fmt.Println(barSUm)
}

func foo(num int) int {
	return num
}

func bar(num int, val string) (int, string) {
	return num, val
}

func foo1(num ...int) int {
	sum := 0
	for _, value := range num {
		sum = sum + value
	}
	return sum

}

func bar1(num []int) int {
	sum := 0
	for _, value := range num {
		sum = sum + value
	}
	return sum

}

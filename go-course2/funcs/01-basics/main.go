package main

import "fmt"

type person struct {
	firstname string
}

func main() {

	fmt.Println(getDogYears("Yussuf", 20))
	s1, _ := dogYears("Todd", 40)
	fmt.Println(s1)
	fmt.Println(dogYears("Todd", 40))
	// slice x
	// running variadic func
	x := []float64{12.50, 7.5, 6.0, 5.3, 8.7, 2.05, 3.95, 10.5, 9.5}
	// Unfurling slice x
	sum1 := sum(x...)
	fmt.Println(sum1)

	// using slice of int
	sum2 := sum(12.09, 0.01)
	fmt.Println(sum2)

	// defer function
	defer testTwo()
	testOne()

	p1 := person{
		firstname: "Yussuf",
	}
	p1.speak()

}

// Attaching method/function to a Type person
func (p person) speak() {
	fmt.Println("I am", p.firstname)
}

// Variadic func
// A variadic function is a function that accepts zero, one, or more values as a single argument.
// For example, fmt.Println is a common variadic function
func sum(ii ...float64) float64 {
	var result float64
	for _, value := range ii {
		result += value
	}
	return result
}
func getDogYears(name string, age int) (string, int) {
	return name + " is now", 7 * age // return multiple values from a function.
}
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

func testOne() {
	fmt.Println("Testing defer function - one")

}
func testTwo() {
	fmt.Println("Testing defer function - two")

}

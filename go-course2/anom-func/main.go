package main

import (
	"fmt"
)

func main() {
	foo()
	// Anonymous functiona
	func() {
		fmt.Println("this is a anonymous func")
	}()
	func(s string) {
		fmt.Println("this is another anom func with my name", s)
	}("Yussuf Ali")

	// funtion expression
	x := func() {
		fmt.Println("This is a func expression example!")
	}
	y := func(age int) {
		fmt.Println("this is another func expression, with my age", age)
	}
	x()
	y(29)

	z := bar()
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", z)
	fmt.Println(z())

}

// returning a func
func bar() func() int {
	return func() int {
		return 2023
	}
}

// Normal func
func foo() {
	fmt.Println("this is an normal func")
}

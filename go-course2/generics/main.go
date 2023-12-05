package main

import (
	"fmt"
)

func addI(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 {
	return a + b
}

// ~ accepts values of types (int,float) and any other value similar underlying types
type myNumbers interface {
	~int | ~float64
}

// similar to interface above but using constraints
// type myNumbers interface {
// 	constraints.Integer | constraints.Float
// }

// generic func accepts int or float64 types - myNumbers
func addG[T myNumbers](a, b T) T {
	return a + b
}

// type alias and underlying type contraints(~)
type myAlias int

func main() {

	var n myAlias = 11
	// only add int func
	fmt.Println(addI(12, n))
	// only add float64 func
	fmt.Println(addF(12.927663, 7.0811022))
	// generic func, works with both int and float64
	fmt.Println(addG(12.55, 8))

}

package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	var one triangle
	one.base = 34.8
	one.height = 5.04

	var two square
	two.sideLength = 10.5

	// fmt.Println(one.getArea())
	printArea(one)
	// fmt.Println(two.getArea())
	printArea(two)

}

func printArea(p shape) {
	fmt.Println(p.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

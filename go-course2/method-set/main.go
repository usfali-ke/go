package main

import "fmt"

// The method set of a type *T consists of all methods with receiver *T or T
// The method set of a type T consists of all methods with receiver type T.

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking!")
}

func (d *dog) run() {
	d.first = "Cheetah"
	fmt.Println("My name is", d.first, "and I'm running!")
}

type young interface {
	walk()
	run()
}

func youngFunc(y young) {
	y.run()
}
func youngFunc2(y young) {
	y.walk()
}

func main() {
	d1 := dog{
		first: "Simba",
	}
	d1.walk()
	d1.run()
	// youngFunc(d1)
	// youngFunc2(d1)

	d2 := &dog{"Bruce"}
	d2.walk()
	d2.run()
	youngFunc(d2)
	youngFunc2(d2)

}

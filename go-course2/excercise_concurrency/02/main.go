package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		name: "Yussuf Ali",
	}
	saySomething(&p1)
	// saySomething(p1)

}

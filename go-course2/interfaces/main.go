package main

import "fmt"

type person struct {
	first string
	last  string
}

type superAgent struct {
	person
	ltk      bool
	agentNum int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last)
}

func (s superAgent) speak() {
	fmt.Println("My name is", s.first, s.last, "and I'm super Agent with LTK", s.ltk)
}

type human interface {
	speak()
}

func saySom(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "Yussuf",
		last:  "Ali",
	}
	p1.speak()

	sa1 := superAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk:      true,
		agentNum: 007,
	}
	sa1.speak()

	saySom(sa1)
	saySom(p1)
}

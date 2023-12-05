package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Embedding struct
type superAgent struct {
	person
	ltk      bool
	agentNum string
}

func main() {
	// #1
	p1 := person{
		first: "Yussuf",
		last:  "Ali",
		age:   32,
	}
	// #2
	p3 := person{"Jonathan", "Kanether", 67}

	s1 := superAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   54,
		},
		ltk:      true,
		agentNum: "sa007",
	}
	// #3
	// Anonymous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Alien",
		last:  "Hacker",
		age:   44,
	}
	fmt.Printf("%T", p1)
	fmt.Println(p1)
	fmt.Printf("%T", p3)
	fmt.Println(p3)
	fmt.Printf("%T", p2)
	fmt.Println(p2)
	fmt.Printf("%T", s1)
	fmt.Println(s1)
}

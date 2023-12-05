package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

func main() {
	i := []int{12, 89, 1, 78, 12, 5, 8, 28, 19, 41}
	s := []string{"test", "string", "respect", "you", "go", "people"}
	fmt.Println("Unsorted data...")
	fmt.Println(i)
	fmt.Println(s)

	fmt.Println("Sorted data...")
	sort.Ints(i)
	fmt.Println(i)

	sort.Strings(s)
	fmt.Println(s)

	p1 := person{"James", 45}
	p2 := person{"Moody", 78}
	p3 := person{"Kakaku", 23}
	p4 := person{"Abdikadir", 12}

	fmt.Println()
	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
}

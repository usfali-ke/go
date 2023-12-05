package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Address string
}

func main() {

	p1 := person{
		First:   "Yussuf",
		Last:    "Warriah",
		Age:     29,
		Address: "P.O Box 28-60300",
	}
	p2 := person{
		First:   "James",
		Last:    "Bond",
		Age:     46,
		Address: "P.O Box 74544545-00100",
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(string(bs))

}

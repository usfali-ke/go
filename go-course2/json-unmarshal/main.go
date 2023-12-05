package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string `json:"First"`
	Last    string `json:"Last"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}

func main() {

	s := `[{"First":"Yussuf","Last":"Warriah","Age":29,"Address":"P.O Box 28-60300"},{"First":"James","Last":"Bond","Age":46,"Address":"P.O Box 74544545-00100"}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)

	var people []person
	// similar as above
	// people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(people)
	for i, v := range people {
		fmt.Println("\nPerson number", i+1)
		fmt.Print(v)
	}
	fmt.Println()

}

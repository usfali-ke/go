package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	name string
}

func (b book) String() string {
	// return "The title of the book is " + b.name)
	return fmt.Sprint("The title of the book is ", b.name)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The count is ", strconv.Itoa(int(c)))
}

// Wrap function fot logging
func logWrite(s fmt.Stringer) {
	log.Println("Error from root:", s.String())
}

func main() {
	bk := book{
		name: "Outliers",
	}
	fmt.Println(bk)

	var n count = 45

	fmt.Println(n)
	log.Println(n)
	logWrite(n)

}

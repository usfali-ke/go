package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Additional info passed!!")
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errorvar = errors.New("Added to errors using var")

func main() {
	f, err := sqrt(-45)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f)

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errorvar
	}
	return math.Sqrt(f), nil
}

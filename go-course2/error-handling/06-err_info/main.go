package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	f, err := sqrt(16)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f)

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Added text to the error")
	}
	return math.Sqrt(f), nil
}

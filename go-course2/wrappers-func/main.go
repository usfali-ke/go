package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("Error in main func: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		// log.Fatalf("Error: %s", err)
		return nil, fmt.Errorf("Eror in a readFile func %s", err)
	}

	return xb, nil

}

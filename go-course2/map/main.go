package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user for a file path.
	fmt.Print("Enter the file path: ")
	var filePath string
	fmt.Scanln(&filePath)

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	// Initialize a map to store word frequencies.
	wordMap := make(map[string]int)

	// Iterate through each line and count words.
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Split the line into words

		for _, word := range words {
			// Remove punctuation and convert to lowercase for consistency.
			cleanedWord := strings.ToLower(strings.Trim(word, ".,!?"))

			// Increment the word count in the map.
			wordMap[cleanedWord]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Print the word frequencies.
	for word, count := range wordMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}

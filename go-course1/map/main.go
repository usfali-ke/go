package main

import (
	"fmt"
)

func main() {
	// One way to create a map
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0001",
		"blue":  "#ff0002",
	}
	fmt.Println(color)
	// Other ways of creating a map
	var color2 map[string]string
	fmt.Println(color2)

	// Another/second way to create a map
	color3 := make(map[string]string)
	fmt.Println(color3)
	// Add to map
	color3["yellow"] = "#fff8999"
	color3["pink"] = "#fff000pp"
	color3["purple"] = "#fff90900"
	color3["orange"] = "#fffor88"
	color3["indigo"] = "#fffind999"

	fmt.Println(color3)
	// Get a item in a map
	fmt.Println(color3["indigo"])

	// Comma ok Idiom
	if v, ok := color3["purple"]; !ok {
		fmt.Println("Doesn't not Exist")
	} else {
		fmt.Println("Does Exist", v)
	}

	// Delete an item
	delete(color3, "pink")
	fmt.Println(color3)

	printMap(color3)
	printMapValue(color3)
}

// Print key and value
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color", color, "is", hex)
	}
}

// Print only value/color
func printMapValue(c map[string]string) {
	for _, hex := range c {
		fmt.Println("color is", hex)
	}
}

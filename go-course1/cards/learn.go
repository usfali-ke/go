package main

import (
	"fmt"
)

func ifElse() string {
	// Go if-else
	var testVal int = 13
	/* check the boolean condition */
	if testVal%2 == 0 {
		/* if condition is true then print the following */
		// fmt.Printf("testVal is even\n")
		// result := strconv.Itoa(testVal) //Convert Int to Str
		return "testVal is even\n"
	} else {
		/* if condition is false then print the following */
		// fmt.Printf("testVal is odd\n")
		return "testVal is odd\n"
	}
}

func ifElse2() string {
	fmt.Print("Enter text: ")
	var input int
	fmt.Scanln(&input)
	if input < 0 || input > 100 {
		// fmt.Print("Please enter valid no")
		return "Please enter valid no"
	} else if input >= 0 && input < 50 {
		// fmt.Print(" Fail")
		return " Fail"
	} else if input >= 50 && input < 60 {
		// fmt.Print(" D Grade")
		return " D Grade"
	} else if input >= 60 && input < 70 {
		// fmt.Print(" C Grade")
		return " C Grade"
	} else if input >= 70 && input < 80 {
		// fmt.Print(" B Grade")
		return " B Grade"
	} else if input >= 80 && input < 90 {
		// fmt.Print(" A Grade")
		return " A Grade"
	} else if input >= 90 && input <= 100 {
		// fmt.Print(" A+ Grade")
		return " A+ Grade"
	} else {
		return "Nothing!!"
	}
}

func nestedIf() string {
	var x int = 10
	var y int = 20
	/* check the boolean condition */
	if x >= 10 {
		/* if condition is true then check the following */
		if y >= 10 {
			/* if condition is true then print the following */
			return "Inside nested If Statement \n"
		}
	}
	return "Nested if-statement"
}

func switchGo() string {
	fmt.Print("Enter Number: ")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 10:
		// fmt.Print("the value is 10")
		return "the value is 10"
	case 20:
		// fmt.Print("the value is 20")
		return "the value is 20"
	case 30:
		// fmt.Print("the value is 30")
		return "the value is 30"
	case 40:
		// fmt.Print("the value is 40")
		return "the value is 40"
	default:
		// fmt.Print(" It is not 10,20,30,40")
		return "Switch Go - It is not 10,20,30,40"
	}
}

func nestedForLoop() (int, int) {
	for a := 0; a < 3; a++ {
		for b := 3; b > 0; b-- {
			return a, b
		}
	}
	return 0, 0
}

// func using args**
func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0
	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

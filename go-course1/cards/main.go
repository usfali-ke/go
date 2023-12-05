package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	// Data types
	var i int
	var f float64
	var bol bool
	var s string
	fmt.Printf("%T %T %T %T\n", i, f, bol, s)              // Prints type of the variable
	fmt.Printf("%v   %v      %v  %q     \n", i, f, bol, s) //prints initial value of the variable

	// Go if-else
	fmt.Println(ifElse())
	fmt.Println(ifElse2())
	// Nested ifstatement
	fmt.Println(nestedIf())
	// Switch
	fmt.Println(switchGo())
	// nested for loop
	fmt.Println(nestedForLoop())

	// func
	fmt.Println(addAll(10, 15, 20, 25, 30))

	// A closure is a function which refers reference variable from outside its body. The function may access and assign to the referenced variables.
	number := 10
	squareNum := func() int {
		number *= number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())

	cards := deck{"Ace of Spades", "One of Hearts"}
	fmt.Println(cards)
	fmt.Println([]byte(cards.toString()))

	cards.print()

	// cards.saveToFile("/home/alien/testfile.txt")
	cards5 := newDeckFromFile("/home/alien/testfile.txt")
	cards5.print()

	cards6 := newDeck()
	cards6.print()
	cards6.shuffle()
	cards6.print()

	/* local variable definition */
	var a int = 378
	var b int = 200
	var ret int
	/* calling a function to get max value */
	ret = max(a, b)
	fmt.Printf("Max value is : %d\n", ret)

	// SecureRandom
	RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	fmt.Println(RandomCrypto)

	cards7 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range cards7 {
		if value%2 == 0 {
			fmt.Println(value, "even")
		} else {
			fmt.Println(value, "odd")
		}
	}
}

func newDesk() deck {
	cards := deck{"Ace of Spades", "One of Hearts"}
	return cards
}

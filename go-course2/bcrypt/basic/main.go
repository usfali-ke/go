package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	spass := `Password@123`
	bs, err := bcrypt.GenerateFromPassword([]byte(spass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(spass)
	fmt.Println(bs)
	fmt.Println(string(bs))

	spass2 := `Password@123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(spass2))
	if err != nil {
		fmt.Println("Password entered is", spass2)
		fmt.Println("Auth Failed")
		return
	}
	fmt.Println("Password entered is", spass2)
	fmt.Println("Auth successful")
}

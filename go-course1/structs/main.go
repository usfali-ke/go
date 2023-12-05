package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //Embedded struct
	// contactInfo		//contact can be defined as just contactInfo, which stands for both var name and var type
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Kamau"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Partey",
		contact: contactInfo{
			email:   "jpartey@gmail.com",
			zipCode: 00200,
		},
	}
	// Pointer to person
	// jimPointer := &jim
	// jimPointer.updatedName("jimmy")
	// OR
	// type of erson
	jim.updatedName("Jimmy")
	jim.print()

}
func (pointerToPersion *person) updatedName(newFirstName string) {
	(*pointerToPersion).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v\n", p)

}

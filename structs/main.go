package main

import "fmt"

// struct for contact information
type contactInfo struct {
	email   string
	zipCode int
}

// struct for person
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// this is the more official way of defining a new struct, relying on order to define variables
	// alex := person{"Alex", "Anderson"}

	// additional way of defining a new struct, this defines a new struct with a Zero Value
	// [String "", int 0, float 0, bool false]
	// var alex person

	//other way of defining a new struct
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	// updating fields on structs
	// alex.firstName = "adam"

	// print-line as is, just curly braces filled with values of struct
	// fmt.Println(alex)

	//print with property names
	fmt.Printf("%+v", alex)

	//definining a new struct that has an embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim

	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// setting the receiver function for the struct
func (p person) print() {
	fmt.Printf("%+v", p)

}

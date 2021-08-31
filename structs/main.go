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

	// the & is an operator and a variable name together; (look at this variable and give us access to the memory address for this variable)
	// jimPointer := &jim

	// we call the update on the pointer which will update the actual value
	// jimPointer.updateName("jimmy")

	// this shorthand will work if we define the reciever to use a pointer Go will allow us to call it on the struct 'person'
	jim.updateName("tex")
	jim.print()
}

// the * value gives us the value that is currently at the memory address provided
// receiver type is *person which means a pointer that points at a person (type description which is saying only a pointer that points to a person)
func (pointerToPerson *person) updateName(newFirstName string) {
	// the *pointerToPerson is an operator it means we want to maniuplate the value that the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

// setting the receiver function for the struct
func (p person) print() {
	fmt.Printf("%+v", p)

}

// Value Types (Use Pointers to Change) : INT , FLOAT , STRING , BOOL , STRUCTS
// Reference Types (Dont use Pointers to Change) : SLICES, MAPS, CHANNELS, POINTERS, FUNCTIONS

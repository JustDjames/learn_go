package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}

type person struct {
	firstName string
	lastName  string
	// uses contactInfo struct
	contactInfo
}

func main() {
	// you can also do
	/* var alex person
	   alex.firstName = "Alex"
	   alex.lastName = "Anderson"
	*/
	bob := person{
		firstName: "Bob",
		lastName:  "Dillian",
		contactInfo: contactInfo{
			email:    "bdillian@aol.com",
			postCode: "n19 8fw",
		},
	}

	// bobPointer gets the person bob variable from memory. if we didn't do this, the actual bob variable wouldn't get updated. instead a copy of the bob variable that Go makes when using calling a function on a variable would get updated
	// The &bob gets the memory address of the value the variable is pointing at
	// for example: the value of the bob variable is at address 0001 in memory. bob points to the value of the bob variable. bobPointer points to the address 0001 in memory
	// pretty much: turn value into address with &value
	bobPointer := &bob
	bobPointer.updateName("Dan")
	bob.print()
}

// *person (or any *<Data Type>) is a  type description. It means we are working with a pointer to a person (or a pointer to whatever data type)
// setting up the function this way means that we don't have to get the pointer as we did on line 36. we could just use bob.updateName("Dan")
func (pointerToPerson *person) updateName(newFirstName string) {
	// the *pointerToPerson  gets the value the memory address is pointing at
	// for example: if bobPointer contains the address in memory of the bob variable, *pointerToPerson will contain the value at that memory address
	// pretty much: turn address into value with *address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// prints key and value of person
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	address struct {
		street string
		city   string
	}
}

func main() {
	// alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.contactInfo.zipCode = 12345
	alex.contactInfo.email = "none"

	jim := person{firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.address.city = "somewhere"

	// jimPointer := &jim

	// jimPointer.updateName("jimmy")
	name := "tester"
	fmt.Println(&name)
	jim.updateName("second")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// var alex person
	// // alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex.firstName = "Subaru"
	// alex.lastName = "just a name"
	// // fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 12102,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName // when we call p is just copy jim to other address in ram so we need to use pointer to fix that
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

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
	// first way of defining a struct
	// alex := person{firstName: "Alex", lastName: "Anderson"} // alex will be first name and anderson will be last name
	// fmt.Println(alex)

	// another way of defining a struct
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) //%+ will print out all the different values from alex
	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 2020,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jimPointer.print()

	//fmt.Printf("%+v", jim)
	//jim.updateName("jimmy")
	//jim.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = "alex"
// }
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

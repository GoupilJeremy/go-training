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
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var jeremy person

	jeremy.firstName = "Jeremy"
	jeremy.lastName = "Goupil"

	fmt.Println(jeremy)
	fmt.Printf("%+v", jeremy)

	jim := person{
		firstName: "Jim",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jim.doe@go.fr",
			zipCode: 91210,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.jim@gmail.com",
			zipCode: "560035",
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	fmt.Println(jim)
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

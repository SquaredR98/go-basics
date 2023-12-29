package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Have some issues with assignment
	alex := person{"Alex", "Aderson", contactInfo{"ravi@mail.com", 900900}}
	// We can avoid any error due to change of order mistakenly
	alex = person{firstName: "Ravi", lastName: "Ranjan"}
	fmt.Println(alex)

	// Another way to define and declare structs
	var sam person
	sam.firstName = "Samuel"
	sam.lastName = "Anderson"
	fmt.Println(sam)

	// Another way of printing struct, i.e., log with the key and value
	fmt.Printf("%+v", sam)

	sam.contact = contactInfo{
		email:   "mail@null.com",
		zipcode: 909878,
	}

	sam.print()

	sam.updateName("Henderson")
	sam.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	p.firstName = newName
}

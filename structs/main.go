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
	arthur := person{
		firstName: "Arthur",
		lastName:  "Dent",
		contactInfo: contactInfo{
			email:   "arthur_dent@example.com",
			zipCode: 12345,
		},
	}
	arthur.updateFirstName("Harvey")
	arthur.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

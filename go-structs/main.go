package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	id        int
	active    bool
	contact   contact
}

type contact struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func modify(p *person) {
	fmt.Println(&p)
	(*p).lastName = "Nathan"
}

func main() {
	rajaram := person{firstName: "Rajaram", lastName: "Padmanathan", active: true, contact: contact{email: "abc@rr.com", zip: 95131}}
	rajaramPtr := &rajaram
	fmt.Println(rajaramPtr)
	(&rajaram).updateName("Raja")
	rajaram.print()
	rajaram.updateName("Ram")
	rajaram.print()
	fmt.Println(&rajaram)
	modify(rajaramPtr)
	rajaram.print()
}

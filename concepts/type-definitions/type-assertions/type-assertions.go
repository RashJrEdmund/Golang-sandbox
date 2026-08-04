/*
X POST
https://x.com/orashus/status/2077069684667543793
*/

package main

import "fmt"

/*
https://go.dev/tour/methods/15
A type assertion provides access to an interface value's
underlying concrete value.
*/

func main() {
	person := Person{name: "Rash", age: 110}

	saySomething(person)
}

type Human interface {
	sayHello(recipient string)
}

type Person struct {
	age                int
	name, phone, email string
}

func (p Person) sayHello(recipient string) {
	fmt.Printf("Hello %s\n", recipient)
}

func (p Person) getNameAndAge() (string, int) {
	return p.name, p.age
}

func saySomething(h Human) {
	h.sayHello("World")

	p, isPerson := h.(Person) // isPerson is a boolean that is true if h is a Person

	if isPerson {
		name, age := p.getNameAndAge()

		fmt.Printf("My name is %s and I am %d years old\n", name, age)
	}
}

/*
DOCS: https://go.dev/blog/maps

The zero value of a map is nil.

We can create a map by using the make() function:
	ages := make(map[string]int)
	ages["John"] = 37

or by using a map literal:
	ages := map[string]int{
		"John": 37,
		"Jane": 32,
	}

We can also use structs as map values
	type car struct {
		registration string
		model        string
	}

	cars := map[string]car{
		"ABC-123": {registration: "ABC-123", model: "Civic"},
	}

We can use the len() function to get the number of key-value pairs in a map
we can use the delete() function to delete a key-value pair from a map
	like: delete(ages, "John")

we can also use the range keyword to iterate over a map
*/

/*
PRINTING FORMATTING:
%+v - Prints the map in a human-readable format
%#v - Prints the map in a Go-syntax format
%T - Prints the type of the map
%p - Prints the address of the map
%#p - Prints the address of the map in hexadecimal
%#x - Prints the address of the map in hexadecimal
%#X - Prints the address of the map in hexadecimal
%#U - Prints the address of the map in hexadecimal
*/

package main

import "fmt"

func main() {
	ages := map[string]int{
		"John": 37,
		"Jane": 32,
	}

	ages["Nathan"] = 30

	delete(ages, "Jane")

	fmt.Printf("Size: %d\n", len(ages))

	fmt.Printf("Ages: %T %p %+v\n", ages, &ages, ages)

	withMake()
}

func withMake() {
	phoneNumbers := make(map[string]int)
	phoneNumbers["John"] = 1234567890
	phoneNumbers["Jane"] = 9876543210
	phoneNumbers["Nathan"] = 1112223333

	delete(phoneNumbers, "Jane")

	fmt.Printf("Size: %d\n", len(phoneNumbers))

	fmt.Printf("Phone Numbers: %T %p %+v\n", phoneNumbers, &phoneNumbers, phoneNumbers)
}

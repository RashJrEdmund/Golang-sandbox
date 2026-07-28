/*
DOCS: https://go.dev/blog/maps

Map types are reference types, like pointers or slices, and so their zero value is nil.

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

/*
You can safely check if a key exists in a map by using the comma ok idiom
	like: age, ok := ages["John"]

	if ok {
		fmt.Printf("Age: %d\n", age)
	} else {
		fmt.Printf("Age not found\n")
	}
*/

package main

import (
	"fmt"
	"maps"
	"slices"
)

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

	fmt.Println("--------------------------------")

	gettingMapKeys()
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

func gettingMapKeys() {
	// OPTIONS A: Using the map.keys()
	/*
		The standard library includes the maps package.
		You can extract keys in a single line using maps.Keys() combined with slices.Collect() to turn the iterator into a slice.
	*/

	userRoles := map[string]string{
		"Alice": "Admin",
		"Bob":   "Editor",
		"Jon":   "Subscriber",
	}

	// 1. Extract keys directly into a slice
	mapKeys := slices.Collect(maps.Keys(userRoles))

	fmt.Println("keys with maps.Keys() and slices.Collect(): ", mapKeys)
	// Note: Order is randomized!

	// OPTION B: By iterating over the map
	/*
		If you are working on an older codebase, or if you want to avoid pulling in packages,
		you write a standard for range loop.

		To optimize performance, always pre-allocate the slice capacity using make([]Type, 0, len(myMap)).
		This prevents Go from resizing the slice in memory during the loop.
	*/

	loopKeys := make([]string, 0, len(userRoles))

	// 2. Loop and grab only the key (leave out the value)
	for key := range userRoles {
		loopKeys = append(loopKeys, key)
	}

	fmt.Println("keys with loop: ", loopKeys)
}

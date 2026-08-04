/*
https://go.dev/ref/spec#Iota

Go has a language feature, that when used with a type definition (and if you squint really hard), kinda looks like an enum (but it's not).
It's called iota.

Within a constant declaration, the pre-declared identifier `iota` represents successive untyped integer constants.
Its value is the index of the respective `ConstSpec` in that constant declaration, starting at zero.
It can be used to construct a set of related constants:

	type sendingChannel int

	const (
    Email sendingChannel = iota
    SMS
    Phone
	)

	The iota keyword is a special keyword in Go that creates a sequence of numbers.
	It starts at 0 and increments by 1 for each constant in the const block.
	So in the example above, Email is 0, SMS is 1, and Phone is 2.

	Go developers sometimes use iota to create a sequence of constants to represent a set of related values,
	much like you would with an enum in other languages.

	But remember, it's not an enum. It's just a sequence of numbers.
*/

package main

import "fmt"

func main() {
	const (
		Apple  = iota // is 0 - which is the index of the first constant
		Banana        // is 1 - which is the index of the second constant
		Cherry        // is 2 - which is the index of the third constant
	)

	fmt.Println("Apple: ", Apple)
	fmt.Println("Banana: ", Banana)
	fmt.Println("Cherry: ", Cherry)
	// --------------------------------------------

	const x = iota // x == 0
	const y = iota // y == 0

	fmt.Println("\nx: ", x)
	fmt.Println("y: ", y)

	// --------------------------------------------
	const (
		a = 1 << iota // a == 1  (iota == 0)
		b = 1 << iota // b == 2  (iota == 1)
		c = 3         // c == 3  (iota == 2, unused)
		d = 1 << iota // d == 8  (iota == 3)
	)

	fmt.Println("\na: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	// --------------------------------------------

	const (
		Sunday  = 1 << iota
		Monday  = 2 << iota
		Tuesday = 3 << iota
		Wednesday
		Thursday
		Friday
		Partyday
		numberOfDays // this constant is not exported
	)

	fmt.Println("\nSunday: ", Sunday)
	fmt.Println("Monday: ", Monday)
	fmt.Println("Tuesday: ", Tuesday)
	fmt.Println("Wednesday: ", Wednesday)
	fmt.Println("Thursday: ", Thursday)
	fmt.Println("Friday: ", Friday)
	fmt.Println("Partyday: ", Partyday)
	fmt.Println("numberOfDays: ", numberOfDays)

}

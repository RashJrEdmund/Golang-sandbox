package main

import "fmt"

func add(a, b int) (c int) {
	c = a + b
	return
}

func nakedReturn(x int) (a, b int) {
	a = x + 1
	b = x + 2
	return
}

func main() {
	c := add(1, 2)

	a, b := nakedReturn(1)

	fmt.Println("Hello from main", c, a, b)
}

// type Person struct {
// 	Name string
// }

// func (p Person) Speak() {
// 	fmt.Println("Hi,", p.Name)
// }

/*
DEFER:
	A defer statement defers the execution of a function until the surrounding function returns.
	this function is evaluated immediately, but not executed until the surrounding function returns.
	Arguments passed to a deferred function are evaluated immediately when the defer line is encountered, not at the end of the function

	func main() {
		i := 0
		defer fmt.Println("Deferred value:", i) // i is evaluated as 0 right here
		i++
		fmt.Println("Current value:", i)
	}
	// Output:
	// Current value: 1
	// Deferred value: 0

In Go, multiple defer statements are executed in Last-In, First-Out (LIFO) order.
When a surrounding function returns, the runtime processes the deferred calls by popping them off a stack,
meaning the last deferred function executes first.

package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("First Defer")
	defer fmt.Println("Second Defer")
	defer fmt.Println("Third Defer")

	fmt.Println("End")
}

OUTPUT:
	Start
	End
	Third Defer
	Second Defer
	First Defer


IMPORTANT:
	Defer statements are executed even if a runtime panic occurs.
	Defer statements are executed even if a surrounding function returns due to a return statement.
	Defer statements are executed even if a surrounding function returns due to a panic.
	Defer statements are executed even if a surrounding function returns due to a panic.
*/

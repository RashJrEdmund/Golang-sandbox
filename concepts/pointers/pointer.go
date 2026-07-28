/*
X POST: https://x.com/orashus/status/2081988664091930964
*/

/*
	A variable is a named location in memory that stores a value.
	We can manipulate the value of a variable by assigning a new value to it or by performing operations on it.
	When we assign a value to a variable, we are storing that value in a specific location in memory
		c := 35
		// "c" is the name of a location in memory, and that locations is storing the integer value 35

	---------------------------------------------------------------------------
	A POINTER IS A VARIABLE THAT STORES THE MEMORY ADDRESS OF ANOTHER VARIABLE
		The zero value of a pointer is nil. Empty pointers are also called "nil pointers"
	---------------------------------------------------------------------------

-->The Asterisk (*) is used to declare a pointer variable type.
		var p *int
		// "p" is a pointer to an integer variable
-->The Asterisk (*) is also used to dereference a pointer variable to get the original value.
		fmt.Println(*p)
		// "*" is used to dereference the pointer variable and print the value stored at the memory address it points to
-->The & operator generates a pointer to it's operand.
		x := 60
		p := &x
		// "p" is now a pointer to the memory address of "x"
*/

/*
	QUICK CHEAT SHEET:
	var p *int   // pointer type
	p = &x       // address of x
	fmt.Println(*p) // dereference
	*p = 42      // modify value through pointer
*/

package main

import "fmt"

func main() {
	x := 42

	p := &x // & -> address of x

	fmt.Println(p)  // 0xc0000120a0 (an address)
	fmt.Println(*p) // * -> 42 (dereferences p to get the value of x)

	*p = 100 // modifying the value of x through the pointer

	fmt.Println(x) // 100

	modifyValueAgain(p) // same as: modifyValueAgain(&x)

	fmt.Println(x) // 200
}

func modifyValueAgain(p *int) {
	*p = 200
}

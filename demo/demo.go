package main

import "fmt"

/**
* Statically typed language.
* data types
* int which has different sizes like int8, int16, int32, int64.
 */

func main() {
	fmt.Println("\nHello, World!")
	// This is explicit assignment, where the variable is declared, and the data type is specified.
	// also explicit assignment with the "var" keyword allows you to modify the value later in the program.
	var age int8 = 20

	fmt.Println("age", age)

	// Reading the data type of the variable.

	const number int32 = 1000000000

	fmt.Printf("number is of type %T\n", number)

	// The implicit assignment involves using the ":=" operator to quickly define and assign a value to a new variable. It also infers the data type of the variable.

	var name = "John"

	// %T is a placeholder for the data type of the variable and shows the data type of the variable..
	fmt.Printf("name is of type %T\n", name)

	// TYPE CASTING
	// Type casting is the process of converting a value from one data type to another.
	// For example, converting an int to a float.

	var num int = 10
	var floatNum float32 = float32(num)

	fmt.Printf("floatNum is of type %T\n", floatNum)

	// Type casting is also used to convert a string to an int.

	var str = "10"
	var intNum int = 10

	fmt.Printf("intNum is of type %T\n", intNum)

	fmt.Println("and here's the str", str)
}

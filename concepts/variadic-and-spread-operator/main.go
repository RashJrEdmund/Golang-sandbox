/*
X POST
https://x.com/orashus/status/2079926280183583104
*/

/*
Functions can take an arbitrary number of final arguments. This is done using the "..." syntax in the function signature.
A variadic function receives the variadic arguments as a slice.

-------
You can also use the spread operator to pass a slice to a variadic function.
*/

package main

import "fmt"

/*
signature for a variadic function with the "..." syntax
arrStr is just a slice of strings
*/
func concat(arrStr ...string) string {
	var res string
	for i := 0; i < len(arrStr); i++ {
		res += arrStr[i]
	}
	return res
}

func main() {
	res1 := concat("Hello ", "there ", "friend!")

	strSlice := []string{"Hi ", "there. ", "Dev over here!"}

	res2 := concat(strSlice...) // using the spread operator "..." to pass a slice to a variadic function
	fmt.Println(res1)
	fmt.Println(res2)

	// Output: Hello there friend!
	// Output: Hi there. Dev over here!
}

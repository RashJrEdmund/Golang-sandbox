/*
Arrays are fixed in size. Once you make an array like [5]int you can't add a 6th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.
see: https://go.dev/tour/moretypes/12

Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
*/
package main

import "fmt"

func main() {
	/*
		primes := [7]int{2, 3, 5, 7, 11, 13, 17}
		mySlice := primes[1:5]
		// mySlice = {3, 5, 7, 11}

		The syntax is:

		arrayName[low:high]
		arrayName[low:]
		arrayName[:high]
		arrayName[:]

		Where low is inclusive and high is exclusive.

		low, high, or both can be omitted to use the entire array on that side of the colon.

		the range is [low, high)
	*/
	primes := [7]int{2, 3, 5, 7, 11, 13, 17}

	mySlice := primes[:1] // prints [2]
	mySlice = primes[4:]  // prints [11, 13, 17]
	mySlice = primes[:]   // prints [2, 3, 5, 7, 11, 13, 17]
	mySlice = primes[2:4] // prints [5, 7]

	fmt.Println(mySlice)
}

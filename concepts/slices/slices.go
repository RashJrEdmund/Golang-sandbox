/*
Arrays are fixed in size. Once you make an array like [5]int you can't add a 6th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.
see: https://go.dev/tour/moretypes/12

Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
*/
package main

import (
	"fmt"
)

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
	mySlice = primes[:]   // prints the whole thing
	mySlice = primes[2:4] // prints [5, 7]
	// mySlice = primes[4:2] // throws an error, invalid slice indices: 2 < 4

	fmt.Println(mySlice)

	fmt.Println("usingAppend(): ", usingAppend())
	fmt.Println("usingMake(): ", usingMake())
	fmt.Println("createMatrix(): ", createMatrix(3, 3))
}

type cost struct {
	day   int
	value float64
}

func usingAppend() []float64 {
	res := []float64{}

	res = append(res, 1.0)
	res = append(res, 2.0)
	res = append(res, 3.0)

	return res
}

func usingMake() []float64 {
	/*
		make is a built-in function that creates a slice.
		It takes the type of the slice, the length of the slice, and the capacity of the slice.
		The capacity is the number of elements the slice can hold.
		The length is the number of elements the slice currently holds.
		The capacity is optional and defaults to the length.
	*/

	res := make([]float64, 3)
	matrix := [][]int{}

	fmt.Printf("matrix: %v\n", matrix)

	res[0] = 1.0
	res[1] = 2.0
	res[2] = 3.0

	return res
}

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		rowVals := []int{}

		for j := 0; j < cols; j++ {
			rowVals = append(rowVals, i*j)
		}

		matrix = append(matrix, rowVals)
	}

	return matrix
}

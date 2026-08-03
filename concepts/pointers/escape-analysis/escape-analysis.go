/*
X POST:
https://x.com/orashus/status/2084364744845045981

Escape analysis in Go is a compile-time process used by the compiler
to determine whether a variable's memory should be allocated on the CPU stack or the managed heap.

The go compiler can even tell you when a variable is moved to heap memory

run:
	go build -gcflags="-m".
*/

/*
MEMORY:
	STACK MEMORY:
		- Extremely fast access (push and pop operations)
		- Self cleaning when function returns
		- Lifetime is tied to and limited to the function scope

	HEAP MEMORY:
		- Slower overhead
		- Cleaned by go's garbage collector
		- Lifetime is Global. Lives until unreferenced.
*/

package main

import "fmt"

func getCount() *int {
	count := 0
	count += 2
	return &count
}

func main() {
	c := getCount() // Now c points to heap memory, which is still valid.
	fmt.Println(*c) // dereferences the pointer and prints the value 2
}

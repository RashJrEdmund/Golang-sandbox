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

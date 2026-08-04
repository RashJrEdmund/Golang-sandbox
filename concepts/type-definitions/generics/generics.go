/*
GENERICS:
Generics help implement reusable, type-safe functions and data structures. DRY code.
	Read on DRY principle: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself

As from go 1.18, support for generics were introduced
Generics allow us to use variables to refer to specific types
*/

/*
Imagine some code that splits a slice into 2 equal parts.
The code that splits the slice doesn't care about the types of values stored in the slice.
Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

	func splitIntSlice(s []int) ([]int, []int) {
		mid := len(s) / 2
		return s[:mid], s[mid:]
	}

	func splitStringSlice(s []string) ([]string, []string) {
		mid := len(s) / 2
		return s[:mid], s[mid:]
	}

	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}

	first, second := splitSlice(intSlice)
	first, second := splitSlice(stringSlice)

But with generics, we can write a single function that can split any slice.

	func splitSlice[T any](s []T) ([]T, []T) {
		mid := len(s) / 2
		return s[:mid], s[mid:]
	}

Now we can call the function with any type of slice.

	first, second := splitSlice(intSlice)
	first, second := splitSlice(stringSlice)

*/

package main

import (
	"fmt"
)

/*
Here any does not work like in typescript.
If a slice of ints is passed, T will simply be int. and not Any like in typescript.
*/
func splitSlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	fmt.Println("mid", mid)
	return s[:mid], s[mid:]
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry", "Pear"}

	inFirst, intSecond := splitSlice(intSlice)
	strFirst, strSecond := splitSlice(stringSlice)

	fmt.Println(inFirst, intSecond)
	fmt.Println(strFirst, strSecond)

	sections()
}

func sections() {
	const separator = "\n--------------------------------"
	fmt.Println(separator)
	fmt.Println("Test Constraints")
	testConstraints(user{name: "John", age: 30})
	fmt.Println(separator)

	fmt.Println("Test Interface Type Lists")
	checkMinimum(1, 2)
	fmt.Println(separator)

	fmt.Println("Test Parametric Constraints")
	parametricConstraints()
	fmt.Println(separator)
}

/*
CONSTRAINTS:
	Constraints are just interfaces that allow us to write generics that only operate within the constraints of a given interface type
*/

type user struct {
	name string
	age  int
}

type User interface {
	GetName() string
	GetAge() int
}

func (u user) GetName() string {
	return u.name
}

func (u user) GetAge() int {
	return u.age
}

func testConstraints[T User](u T) {
	fmt.Printf("Name: %s, Age: %d", u.GetName(), u.GetAge())
}

/*
INTERFACE TYPE LISTS:
	When generics were released, a new way of writing interfaces was also released at the same time!

	Traditional interfaces in Go are method-based: a type satisfies an interface if it has the required methods.

	With generics, we also got type-set (type-list) interfaces. Instead of listing methods, they list the concrete (or underlying) types that are allowed. These are mostly used as constraints on type parameters.

	For example, to use < and > on a type parameter T, the compiler must know T is ordered. A type-list interface spells out exactly which types count as ordered:
*/

// Ordered matches any type that supports <, <=, >, and >=.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Because T is constrained by Ordered, the compiler knows
// that < is valid for any T used with this function.
func checkMinimum[T Ordered](a, b T) {
	if a < b {
		fmt.Printf("%v is less than %v", a, b)
	}
	fmt.Printf("%v is greater than %v", a, b)
}

/*
PARAMETRIC CONSTRAINTS:
	Interface definitions which can be used as constrain can accept type parameters as well.
*/
// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of products the store sells.
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
	booksSold []book
}

// Sell adds a book to the bookStore's sold slice.
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's sold slice.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func parametricConstraints() {
	bs := bookStore{
		booksSold: []book{},
	}

	// By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Printf("\nBooks Sold:\n%+v\n", bs.booksSold)

	// We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts(&ts, []toy{
		{
			name:  "LEGO bricks",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Printf("\nToys Sold:\n%+v\n", ts.toysSold)
}

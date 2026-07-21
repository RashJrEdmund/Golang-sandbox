package main

import "fmt"

type error interface {
	Error() string
}

type someEntity struct {
	key        string
	anotherKey bool
}

func (s someEntity) Error() string {
	return fmt.Sprintf("someEntity: %s, %t", s.key, s.anotherKey)
}

func (s someEntity) GetKey() string {
	return s.key
}

func main() {
	entity := someEntity{
		key:        "test val",
		anotherKey: true,
	}

	isError := isAnError(entity)

	fmt.Println("entity Implements error interface: ", isError) // prints true
}

func isAnError(v error) bool {
	_, isError := v.(error)
	return isError
}

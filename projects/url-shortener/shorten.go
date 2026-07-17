package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println("Hello, World!")

	fmt.Println(args, len(args))
}

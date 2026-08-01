package main

import (
	"fmt"
	"os"
	"strings"

	ms "github.com/RashJrEdmund/go-sandbox/tree/main/bootdev/mystrings"
)

func printHelp() {
	fmt.Println("Usage: reverse <phrase>")
	fmt.Println("Reverse reverses the order of words in a string")
	fmt.Println("Example: reverse 'Hello, World!'")
	fmt.Println("Output: !dlroW ,olleH")
}

func main() {
	if len(os.Args) <= 1 {
		printHelp()
		return
	}

	args := os.Args[1:]

	if strings.ToLower(args[0]) == "--help" || strings.ToLower(args[0]) == "-h" {
		printHelp()
		return
	}

	for _, arg := range args {
		fmt.Println(ms.Reverse(arg))
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unknown operator: %s", op)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./calc <num1> <operator> <num2>")
		os.Exit(1)
	}

	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[3], 64)
	op := os.Args[2]

	result, err := calculate(a, b, op)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}

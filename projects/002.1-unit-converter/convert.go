package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var toMeters = map[string]float64{
	"m":  1.0,
	"km": 1000.0,
	"cm": 0.01,
	"mi": 1609.34,
	"ft": 0.3048,
}

func convert(value float64, from, to string) (float64, error) {
	fromRation, ok := toMeters[from]
	if !ok {
		return 0, fmt.Errorf("Unknown unit: %s", from)
	}
	toRatio, ok := toMeters[to]
	if !ok {
		return 0, fmt.Errorf("Unknown unit: %s", to)
	}
	return value * (fromRation / toRatio), nil
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("usage: ./convert <value> <unit_from> to <unit_to> // example. ./convert 50 m to km")
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("invalid number:", os.Args[1])
		os.Exit(1)
	}

	unitFrom := strings.ToLower(os.Args[2])
	unitTo := strings.ToLower(os.Args[4])

	result, err := convert(value, unitFrom, unitTo)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf("%g %s = %g %s\n", value, unitFrom, result, unitTo)
}

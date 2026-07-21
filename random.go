// USING THIS TO RUN RANDOM CODE AND SEE THE OUTPUT

package main

import "fmt"

func main() {
	const one int = 1

	// for i := 0.0; i < 2; i++ {
	// 	addition := 1.0 + (i / 100)
	// 	fmt.Printf("addition: %f\n %T\n", addition, addition)
	// }

	counter := 0
	for true {
		fmt.Println("true")
		counter++
		if counter > 3 {
			break
		}
	}

	for i := 0; i < len("sd"); i++ {
		fmt.Println("i", i)
	}
	// how
}

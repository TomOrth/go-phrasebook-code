package main

import "fmt"

func main() {
	loops := 1

	// No while keyword
	// Golang way of creating a while loop
	for loops > 0 {
		fmt.Printf("Number of loops: ")

		// Example of user input
		fmt.Scanf("%d\n", &loops)

		// Regular for loop
		for i := 0; i < loops; i += 1 {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
	}
	// Infinite for loop
	for {
		break
	}
}

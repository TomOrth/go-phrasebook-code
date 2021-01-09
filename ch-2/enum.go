package main

import "fmt"

func main() {

	// Enum creation
	const (
		Red = (1<<iota)
		Green = (1<<iota)
		Blue, ColorMask = (1<<iota), (1<<(iota+1))-1
	)

	fmt.Println(Red, Green, Blue, ColorMask)
}

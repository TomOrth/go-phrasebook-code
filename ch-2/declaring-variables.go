package main

import "fmt"

func main() {
	// variables have type after
	var x int = 6
	var y float32 = 3.2

	// Can also use := to infer type
	//chan_name := make(chan interface{})
	fmt.Printf("%d %f\n", x, y)
}

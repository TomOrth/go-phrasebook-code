package main

import "fmt"

type empty interface{}

type example interface {
	notImplemented()
}

func main() {
	one := 1
	// Since empty is of type interface{}, it will take on any type
	var i empty = one

	// Cast from int to float
	var float float32
	float = float32(one)

	// Switch on variable type
	switch i.(type) {
		default:
			fmt.Println("Type Error!")
		case int:
			fmt.Printf("%d\n", i)
	}
	fmt.Printf("%f\n", float)

	// This causes a panic since int does not implement all the methods (this case, 1) of the example interface
	var e example = i.(example)
	fmt.Printf("%d\n", e.(empty).(int))
}
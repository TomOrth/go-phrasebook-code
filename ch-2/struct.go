package main

import "fmt"

// Example Struct
// Uppercase letter makes something public (accessible outside the package)
type Example struct {
	Val string
	count int
}

// Example method for a struct
// Before function name is the receiever, which attaches the method to the struct
func (e *Example) Log() {
	e.count++
	fmt.Printf("%d %s\n", e.count, e.Val)
}

func main() {

	// Example struct creation
	// Since its a pointer, must have &
	x := &Example{
		Val: "Hello",
		count: 2,
	}
	x.Log()
	x.Log()
	fmt.Printf("%+v\n", x)
}

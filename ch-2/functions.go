package main

import "fmt"

// Example function declaration
// Args is an example of variable amount of args
// Returns are declared as a tuple of types
func printf(str string, args ...interface{}) (int, error) {
	_, err := fmt.Printf(str, args...)
	return len(args), err
}

func main() {
	count := 1
	closure := func(msg string) {
		printf("%d %s\n", count, msg)
		count += 1
	}
	closure("Message")
	closure("Another")
}

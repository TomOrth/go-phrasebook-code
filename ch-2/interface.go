package main

// Multiple imports
import (
	"fmt"
	"math"
)

// Different structs and methods that are semi-related

type cartesianPoint struct {
	x, y float64
}

func (p cartesianPoint) X() float64 {
	return p.x
}

func (p cartesianPoint) Y() float64 {
	return p.y
}

func (p cartesianPoint) Print() {
	fmt.Printf("(%f, %f)\n", p.x, p.y)
}

type polarPoint struct {
	r, theta float64
}

func (p polarPoint) X() float64 {
	return p.r * math.Cos(p.theta)
}

func (p polarPoint) Y() float64 {
	return p.r * math.Sin(p.theta)
}

func (p polarPoint) Print() {
	fmt.Printf("(%f, %f)\n", p.r, p.theta)
}

// Example Interface
// First line is another interface (sort of like inheritance)
// The rest define methods related to the interface
type Point interface {
	Printer
	X() float64
	Y() float64
}

type Printer interface {
	Print()
}

// Function takes in an instance of the interface
func demo(point Point) {
	fmt.Printf("%f %f\n", point.X(), point.Y())
	point.Print()
}

func main() {

	// The two structs follow the interface outline
	// Therefore, we can pass it to demo
	cartPoint := cartesianPoint {
		10,
		10,
	}

	polPoint := polarPoint {
		10,
		0.6,
	}
	demo(cartPoint)
	demo(polPoint)
}

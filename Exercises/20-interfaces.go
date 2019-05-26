package main

import "fmt"
import "math"

/*
 Interfaces are named collections of method signatures.
*/

//Here’s a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

//For our example we’ll implement this interface on rect and circle types.
type rect1 struct {
	width, height float64
}
type circle struct {
	radius float64
}

//To implement an interface in Go, we just need to implement all the methods in the interface.
func (r rect1) area() float64 { //Here we implement geometry on rects.
	return r.width * r.height
}
func (r rect1) perim() float64 {
	return (2 * r.width) + (2 * r.height)
}

//The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//If a variable has an interface type, then we can call methods that are in the named interface.
func measure(g geometry) {
	//Here’s a generic measure function taking advantage of this to work on any geometry.
	fmt.Println("Print inside the measure function - g ->", g)
	fmt.Println("Printing the area of geometry - g.area() ->", g.area())
	fmt.Println("Printing the area of perimeter - g.perim() ->", g.perim())

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//The circle and rect struct types both implement the geometry interface
	//so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)

}

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
If a variable has an interface type, then we can call methods that are in the named interface.
Here’s a generic measure function taking advantage of this to work on any geometry.
*/
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Println("\n")
}

func describe(g geometry) {
	fmt.Printf("(%v, %T)\n", g, g)
}

func main() {
	var g geometry
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	describe(g)
	describe(r)
	describe(c)

	measure(r)
	measure(c)
}

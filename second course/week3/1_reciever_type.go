/*
In golang we create methods for OOP using receiver types
A receiver type is a function that has a parathensis containing
an acronym and the name of the type
*/

package main

import (
	"fmt"
	"math"
)

type MyInt int

// below is the receiver type
func (x MyInt) Double() MyInt {
	return x * 2
}

// creating a class with many type of variables
type Point struct {
	x float64
	y float64
}

func (p Point) DistOrigin() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	x := MyInt(2)

	fmt.Println(x.Double())

	p1 := Point{
		x: 3,
		y: 4,
	}

	fmt.Println(p1.DistOrigin())
}

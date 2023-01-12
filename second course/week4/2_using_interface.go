package main

import (
	"fmt"
	"math"
)

/*
Ways to use an interface
1. Need a function which takes multiple types of parameters
2. Empty interface
*/


/* Need a function which takes multiple type of parameters

the example below is used to solve the problem of finding the
best shape for a swimming pool.

there are 2 limitations
1. the area of the compound which the swimming pool will be dug in
2. the length of the fence that will surround the swimming pool
*/

type Shape2d interface {
	area() float64
	perimeter() float64
}

func FitInYard(s Shape2d) bool {
	if s.area() < 100 && s.perimeter() < 100 {
		return true
	}
	return false
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

/* Empty interface specifies no method
- All types satisfy the empty interface
- Use it to have a function accept any type as a parameter
*/

func main() {
	t1 := Triangle{2, 3, 4}
	fmt.Println(FitInYard(t1))
}

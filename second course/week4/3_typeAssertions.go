/*
TYPE ASSERTIONS
- They are used to disambiguate
- They can be used to determine and extract the underlying
concrete type

*/

package main

import (
	"fmt"
	"math"
)

type Shape2d interface {
	area() float64
	perimeter() float64
}

// Type assertion function
func DrawShape(s Shape2d) {
	rect, ok := s.(Rectangle)
	if ok {
		fmt.Println(rect)
	}
	tri, ok1 := s.(Triangle)
	if ok1 {
		fmt.Println(tri)
	}
}

func DrawShape2(s Shape2d) {
	switch sh := s.(type) {
	case Rectangle:
		fmt.Println("A rectangle.", sh)
	case Triangle:
		fmt.Println("A Triable", sh)
	}
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

type Rectangle struct {
	l, b float64
}

func (r Rectangle) area() float64 {
	return r.l * r.b
}

func (r Rectangle) perimeter() float64 {
	return r.l*2 + 2*r.b
}

func main() {
	t1 := Rectangle{3, 4}
	DrawShape(t1)
	DrawShape2(t1)

	t2 := Triangle{3, 4, 6}
	DrawShape(t2)
	DrawShape2(t2)
}

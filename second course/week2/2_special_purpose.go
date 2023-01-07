package main

import (
	"fmt"
	"math"
)

// A function that returns another function

func makeDistOrigin(oX, oY float64) func(x, y float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-oX, 2) + math.Pow(y-oY, 2))
	}

	return fn
}

func main() {
	dist1 := makeDistOrigin(0, 0)
	dist2 := makeDistOrigin(2, 2)

	fmt.Println(dist1(2, 2))
	fmt.Println(dist2(2, 2))
}

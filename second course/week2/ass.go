package main

import (
	"fmt"
)

func main() {
	var acc, iniVelo, iniDisp float64
	fmt.Print("\nEnter accelaration :")
	fmt.Scanf("%f", &acc)
	fmt.Print("\nEnter initial Velocity :")
	fmt.Scanf("%f", &iniVelo)
	fmt.Print("\nEnter initial Displacement :")
	fmt.Scanf("%f", &iniDisp)

	fn := GenDisplaceFn(acc, iniVelo, iniDisp)

	var time float64
	fmt.Print("\nEnter timeTraveled :")
	fmt.Scanf("%f", &time)
	fmt.Print("\nTotal Distance travelled :", fn(time))
	fmt.Print("\n")
}

func GenDisplaceFn(a float64, vot float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5)*t*t*a + (vot * t) + so
	}
	return fn
}

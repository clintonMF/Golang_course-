package main

import "fmt"

func getMax(vars ...int) int {
	maxV := -1

	for _, v := range vars {
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}

func main() {
	fmt.Println(getMax(5, 4, 23, 56, 2, 6))

	// You can also pass in a slice
	sliceV := []int{2, 3, 4, 5, 61, 233, 0, 6}
	fmt.Println(getMax(sliceV...))
}

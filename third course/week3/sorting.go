package main

import (
	"fmt"
)

// This function performs the sorting operation on each array
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	lesser := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range arr[1:] {
		if v <= pivot {
			lesser = append(lesser, v)
		} else {
			greater = append(greater, v)
		}
	}

	lesser_sorted := quickSort(lesser)
	greater_sorted := quickSort(greater)
	lesser_sorted = append(lesser_sorted, pivot)
	lesser_sorted = append(lesser_sorted, greater_sorted...)

	return lesser_sorted
}

/*
This function combines sorted arrays to give one sorted array
*/
func merge(left, right []int) []int {
	sorted := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
	}

	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
	}

	return sorted
}

/*
this function implements a goroutine channel on the sorting function
*/
func routine(arr []int, c chan []int) {
	fmt.Printf("array to be sorted -> %v\n", arr)
	sorted := quickSort(arr)
	c <- sorted
}

func main() {

	// get the user input
	fmt.Println("Please enter a series of integers.\nNumber of integers entered should be greater than 4")
	arr := []int{}
	var i int
	for i > -100 {
		fmt.Println("\nwhen done inputing integers enter a number less than - 100 e.g -150 ")
		fmt.Scan(&i)
		if i > -100 {
			arr = append(arr, i)
		}
	}

	// Create the 4 channels
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)

	// breaking the array into 4
	var frac int
	frac = len(arr) / 4
	arr1 := arr[:frac]
	arr2 := arr[frac : 2*frac]
	arr3 := arr[2*frac : 3*frac]
	arr4 := arr[3*frac:]

	// using goroutine to solve the arrays
	go routine(arr1, c1)
	go routine(arr2, c2)
	go routine(arr3, c3)
	go routine(arr4, c4)

	// receiving the sorted array from the goroutines
	a1 := <-c1
	a2 := <-c2
	a3 := <-c3
	a4 := <-c4

	// merging the sorted arrays
	res1 := merge(a1, a2)
	res2 := merge(a3, a4)
	res3 := merge(res1, res2)

	fmt.Println("\nSorted array")
	fmt.Println(res3)
}

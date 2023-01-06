/* Assingment
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import "fmt"

// func bubbleSort(arr *[]int) {
// 	for i := 0; i < len(*arr)-1; i++ {
// 		Swap(arr, i)
// 	}
// 	fmt.Println("")
// 	fmt.Printf("sorted array => ")
// 	fmt.Println(*arr)
// }

func Swap(arr *[]int, i int) {
	ar := *arr
	for j := 0; j < len(*arr)-i-1; j++ {
		if ar[j] > ar[j+1] {
			ar[j], ar[j+1] = ar[j+1], ar[j]
		}
	}
}

// //This function ask user for input and appends it to the array
// func addInt(arr *[]int) {
// 	var digit int
// 	fmt.Scan(&digit)
// 	*arr = append(*arr, digit)
// }

func main() {
	var arr []int = []int{3, 2, 3, 54, 1, 5, 6, 7}
	// fmt.Println("please enter digits you are to enter 10 digits")
	// i := 0
	// for i < 10 {
	// 	fmt.Println("\nplease enter digit", i+1)
	// 	addInt(&arr)
	// 	i++
	// }

	// bubbleSort(&arr)

	Swap(&arr, 0)
	Swap(&arr, 1)
	Swap(&arr, 2)
	Swap(&arr, 3)
	Swap(&arr, 4)
	Swap(&arr, 5)
	fmt.Println(arr)
}

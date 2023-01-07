package main

import "fmt"

/*
functions in golang are first class functions
This means they can be treated as any other data type
1. variables can be declared with function type
2. Can be created dynamically
3. Can be passed as arguments and returned as values
4. Can be stored in data structures
*/

// Declaring variable as a function

var funcVar func(int) int

func plusOne(a int) int {
	return a + 1
}

// Passing a function as an argument to another function
func applyIt(aFunc func(int) int, x int) int {
	return aFunc(x)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
	// Declaring variable as a function
	funcVar = plusOne
	fmt.Println(funcVar(2))

	// Passing a function as an argument
	fmt.Println("value of incFn:", applyIt(incFn, 5))
	fmt.Println("value of decFn:", applyIt(decFn, 5))

	//Anonymous functions
	v := applyIt(func(x int) int { return x + 1 }, 3)

	fmt.Println("value from the anonymous function", v)

}

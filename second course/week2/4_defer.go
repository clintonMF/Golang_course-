package main

import "fmt"

/*
The defer statement is used before a function to prevent the
function from being executed when they are called. They get
executed later i.e usually at the end of the entire operation.
An example of this is the close function, perhaps while working
on a file you want to close it at the end. you can call the
close function immediately and defer it. hence, it will close
at the end of the operation.
NOTE: the function takes the value of the argument at the point
where it is called.
*/

func welcome() {
	defer fmt.Println("Bye")

	fmt.Println("Hello")
}

func note() {
	i := 1
	defer fmt.Println("i is ", i)
	i++
	fmt.Println("i becomes", i, "but this is printed before the deferred function")
}

func main() {
	welcome()
	note()
}

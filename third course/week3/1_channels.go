package main

import "fmt"

func prod(val1, val2 int, c chan int) {
	c <- val1 * val2
}

func main() {
	var c chan int = make(chan int)

	go prod(3, 4, c)
	go prod(2, 2, c)

	a := <-c
	b := <-c

	fmt.Println(a * b)
}

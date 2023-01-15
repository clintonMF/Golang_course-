package main

import (
	"fmt"
	"regexp"
	"strings"
)

// import (
// 	"fmt"
// 	"time"
// )

// var x int = 4

// func printX() {
// 	fmt.Println(x)

// }

// func incrementX() {
// 	x = x + 1
// }

// func main() {
// 	go incrementX()
// 	go printX()
// 	time.Sleep(10 * time.Second)
// }

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	lower := strings.ToLower(s)
	fmt.Println(lower)

	left := 0
	right := len(lower) - 1

	for left < right {
		if lower[left] != lower[right] {
			return false

		}
		left = left + 1
		right = right - 1
	}

	return true
}

func main() {
	str := "HER i am to say"

	isPalindrome(str)
}

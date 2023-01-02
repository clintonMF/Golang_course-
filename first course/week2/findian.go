/* Assignment 2
Write a program which prompts the user to enter a string. The program
searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the
character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be
case-sensitive, so it does not matter if the characters are upper-case or
lower-case.

Examples: The program should print “Found!” for the following example
entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program
should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program,
“findian.go”. */

package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	fmt.Println("please enter a string")
// 	var inputStr string
// 	fmt.Scan(&inputStr)

// 	// Get the string to lower case to ensure it isn't case sensitive
// 	lowerStr := strings.ToLower(inputStr)

// 	first_letter := lowerStr[0:1]
// 	l := len(lowerStr) //get the length of the string
// 	last_letter := lowerStr[l-1:]
// 	containsA := strings.Contains(lowerStr, "a")
// 	if first_letter == "i" && last_letter == "n" && containsA {
// 		fmt.Println("Found!")
// 	} else {
// 		fmt.Println("Not Found!")

// 	}
// }

// The above was my submitted assignment for which i got 100% as it does the work
// On checking that of fellow students i found out there was a better and easier
// way to implement this.

func main() {
	fmt.Println("Enter a string")
	var inputStr string
	fmt.Scan(&inputStr)

	lower := strings.ToLower(inputStr)
	if strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n") && strings.Contains(lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

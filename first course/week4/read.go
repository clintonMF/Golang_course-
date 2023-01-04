/* Assignment 2
Write a program which reads information from a file and represents it
in a slice of structs. Assume that there is a text file which contains
a series of names. Each line of the text file has a first name and a
last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for
the first name, and lname for the last name. Each field will be a
string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your
program will successively read each line of the text file and create a
struct which contains the first and last names found in the file. Each
struct created will be added to a slice, and after all lines have been
read from the file, your program will have a slice containing one struct
for each line in the file. After reading all lines from the file, your
program should iterate through your slice of structs and print the first
and last names found in each struct. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name_of_file string

	fmt.Println("Please enter the name of the file\nNote: don't forget to add .txt")
	fmt.Scan(&name_of_file)

	type names struct {
		first_name string
		last_name  string
	}

	var names_slice []names

	readFile, err := os.Open(name_of_file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		name := strings.Split(fileScanner.Text(), " ")
		first_name := name[0]
		last_name := name[1]

		each_name := names{
			first_name: first_name,
			last_name:  last_name,
		}

		names_slice = append(names_slice, each_name)

	}
	readFile.Close()

	for i := 0; i < len(names_slice); i++ {
		firstName := names_slice[i].first_name
		lastName := names_slice[i].last_name

		fmt.Printf("First name - %s \nLast name - %s\n\n", firstName, lastName)
	}

}

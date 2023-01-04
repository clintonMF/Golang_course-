/* Assignment 1
Write a program which prompts the user to first enter a name, and
then enter an address. Your program should create a map and add
the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	db := make(map[string]string)
	var name, address string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter your address")
	fmt.Scan(&address)

	db["name"] = name
	db["address"] = address

	fmt.Println("\n\nThe json object is below")
	byte, _ := json.Marshal(db)
	fmt.Println(string(byte))
}

/* ASSINGMENT QUESTION

PROMPT
Write a program which allows the user to get information about a predefined
set of animals. Three animals are predefined, cow, bird, and snake. Each 
animal can eat, move, and speak. The user can issue a request to find out 
one of three things about an animal: 1) the food that it eats, 2) its method 
of locomotion, and 3) the sound it makes when it speaks. The following table 
contains the three animals and their associated data which should be hard-coded 
into your program.

Animal  Food eaten    Locomotion method    Spoken sound

cow       grass            walk               moo

bird		worms			fly				 peep

snake		mice		   slither			  hsss

Your program should present the user with a prompt, “>”, to indicate that the user 
can type a request. Your program accepts one request at a time from the user, prints 
out the answer to the request, and prints out a new prompt. Your program should continue 
in this loop forever. Every request from the user must be a single line containing 2 strings. 
The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string 
is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type 
called Animal which is a struct containing three fields:food, locomotion, and noise, all of 
which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type 
of all of your methods should be your Animal type. The Eat() method should print the animal’s 
food, the Move() method should print the animal’s locomotion, and the Speak() method should print 
the animal’s spoken sound. Your program should call the appropriate method when the user makes 
a request.

*/

package main

import "fmt"

type Animal struct {
	FoodEaten        string
	LocomotionMethod string
	SpeakSound       string
}

func (a Animal) Eat() string {
	return a.FoodEaten
}

func (a Animal) Move() string {
	return a.LocomotionMethod
}
func (a Animal) Speak() string {
	return a.SpeakSound
}

func action(animal Animal, action string) string {
	switch action {
	case "move":
		return animal.Move()
	case "eat":
		return animal.Eat()
	case "speak":
		return animal.Speak()
	default:
		return "Invalid animal action \nTry: move, eat and speak "
	}
}

func main() {
	// cow := Animal{
	// 	FoodEaten:        "grass",
	// 	LocomotionMethod: "walk",
	// 	SpeakSound:       "moo",
	// }
	// bird := Animal{
	// 	FoodEaten:        "worms",
	// 	LocomotionMethod: "fly",
	// 	SpeakSound:       "peep",
	// }

	// snake := Animal{
	// 	FoodEaten:        "mice",
	// 	LocomotionMethod: "slither",
	// 	SpeakSound:       "hsss",
	// }

	animal := map[string]Animal{
		"cow": {"grass", "walk", "moo"},
	}

	s := animal["cow"]

	fmt.Println(s.Eat())
	fmt.Println(s.Speak())
	fmt.Println(s.Move())
	// for {
	// 	var animalName, animalInfo string
	// 	fmt.Println("please enter the name of animal and animal information both in lowercase e.g (cow move)")
	// 	fmt.Scan(&animalName, &animalInfo)
	// 	animalInfo = strings.ToLower(animalInfo)
	// 	animalName = strings.ToLower(animalName)

	// 	if animalName == "cow" {
	// 		fmt.Println(action(cow, animalInfo))
	// 	} else if animalName == "bird" {
	// 		fmt.Println(action(bird, animalInfo))
	// 	} else if animalName == "snake" {
	// 		fmt.Println(action(snake, animalInfo))
	// 	} else {
	// 		fmt.Println("This is not a valid animal name e.g cow, snake, bird")
	// 	}

	// 	fmt.Println()
	// }
}

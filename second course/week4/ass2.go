package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name, food, locomation, noise string
}
type Bird struct {
	name, food, locomation, noise string
}
type Snake struct {
	name, food, locomation, noise string
}

func (cow Cow) Eat() {

	fmt.Printf("eats grass")

}
func (cow Cow) Move() {

	fmt.Printf("walks")

}
func (cow Cow) Speak() {

	fmt.Printf("moo")

}

func (bird Bird) Eat() {

	fmt.Printf("eats worms")

}
func (bird Bird) Move() {

	fmt.Printf("flies")

}
func (bird Bird) Speak() {

	fmt.Printf("peep")

}

func (snake Snake) Eat() {

	fmt.Printf("eats mice")

}
func (snake Snake) Move() {

	fmt.Printf("slither")

}
func (snake Snake) Speak() {

	fmt.Printf("Snake hsss")

}

func main() {

	var requestType, name, animalType, animalInfo string
	var c1 Cow
	var b1 Bird
	var s1 Snake
	var cowSlice = make([]Cow, 0)
	var birdSlice = make([]Bird, 0)
	var snakeSlice = make([]Snake, 0)
	for {
		fmt.Printf("\nEnter the type of request \"newanimal\" or \"query\"\n")
		fmt.Printf(">")
		fmt.Scan(&requestType)
		if requestType == "newanimal" {
			fmt.Printf("\nEnter \"newanimal\" name of the animal and type of animal either \"cow\", \"bird\" or \"snake\" \n")
			fmt.Scan(&requestType, &name, &animalType)
			if animalType == "cow" {
				c1 = Cow{name, "grass", "walk", "moo"}
				cowSlice = append(cowSlice, c1)
			} else if animalType == "bird" {
				b1 = Bird{name, "worms", "fly", "peep"}
				birdSlice = append(birdSlice, b1)
			} else if animalType == "snake" {
				s1 = Snake{name, "mice", "slither", "hsss"}
				snakeSlice = append(snakeSlice, s1)
			}
			fmt.Println("Created it!")
		} else if requestType == "query" {
			fmt.Printf("\nEnter \"query\" name of the animal and information of animal either \"eat\", \"move\" or \"speak\" \n")
			fmt.Scan(&requestType, &name, &animalInfo)
			for _, info := range cowSlice {
				if animalInfo == "eat" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Eat()
					}
				} else if animalInfo == "move" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Move()
					}
				} else if animalInfo == "speak" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Speak()
					}
				}
			}
			for _, info := range birdSlice {
				if animalInfo == "eat" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Eat()
					}
				} else if animalInfo == "move" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Move()
					}
				} else if animalInfo == "speak" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Speak()
					}
				}
			}
			for _, info := range snakeSlice {
				if animalInfo == "eat" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Eat()
					}
				} else if animalInfo == "move" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Move()
					}
				} else if animalInfo == "speak" {
					if info.name == name {
						fmt.Printf(info.name + " ")
						info.Speak()
					}
				}

			}
		}
	}
}

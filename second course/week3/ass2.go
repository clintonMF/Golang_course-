package main

import "fmt"

type Behaviour interface {
	Eat() string
	Move() string
	Speak() string
	initAnimal([]string) *map[string]Animal
}

type Animal struct {
	food, locomotion, spoken string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.spoken
}

func initAnimal(animalMap map[string]Animal) {
	animalMap["cow"] = Animal{"grass", "walk", "moo"}
	animalMap["bird"] = Animal{"worms", "fly", "peep"}
	animalMap["snake"] = Animal{"mice", "slither", "hsss"}
}

var (
	animal = make(map[string]Animal)
)

func main() {
	animalMap := make(map[string]Animal)
	initAnimal(animalMap)
	var (
		animal    = ""
		behaviour = ""
	)

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &behaviour)
		switch behaviour {
		case "eat":
			fmt.Println(animalMap[animal].Eat())
			break
		case "move":
			fmt.Println(animalMap[animal].Move())
			break
		case "speak":
			fmt.Println(animalMap[animal].Speak())
			break
		}
	}
}

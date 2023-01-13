package main

import (
	"fmt"
	"log"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Creature struct {
	name       string
	food       string
	locomotion string
	sound      string
}

type Cow Creature
type Bird Creature
type Snake Creature

func main() {
	fmt.Println("I'm waiting your orders ... valid commands are:')")
	fmt.Println("  - `newanimal` [animal name] [animal type - valid values are `cow`, `bird`, `snake`]')")
	fmt.Println("  - `query` [animal name of already created animal] [information - valid values are `eat`, `move`, `speak`]')")
	fmt.Println()

	var command, arg1, arg2 string
	animals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&command, &arg1, &arg2)
		handleErrorIfAny(err, false)
		switch command {
		case "newanimal":
			animals[arg1], err = newAnimal(arg1, arg2)
			handleErrorIfAny(err, false)
			if err == nil {
				fmt.Println("Created it!")
			}
		case "query":
			if len(animals) == 0 {
				log.Println("Noting to query, you haven't created any animals!")
				continue
			}
			animal, ok := animals[arg1]
			if !ok {
				log.Printf("No animal defined with name: %s\n", arg1)
				continue
			}
			switch arg2 {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				log.Printf("Unknown information: %s, valid values are (eat|move|speak)\n", arg1)
				continue
			}
		default:
			log.Printf("Unknown command: `%s`, valid values are (newanimal|query)\n", command)
			continue
		}
	}

}

func newAnimal(name, animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return NewCow(name), nil
	case "bird":
		return NewBird(name), nil
	case "snake":
		return NewSnake(name), nil
	default:
		return nil, fmt.Errorf("Unknown animal type: `%s`, valid values are (cow|bird|snake)\n", animalType)
	}
}

func NewCow(name string) *Cow {
	return &Cow{
		name:       name,
		food:       "grass",
		locomotion: "walk",
		sound:      "moo",
	}
}

func NewBird(name string) *Bird {
	return &Bird{
		name:       name,
		food:       "worms",
		locomotion: "fly",
		sound:      "peep",
	}
}

func NewSnake(name string) *Snake {
	return &Snake{
		name:       name,
		food:       "mice",
		locomotion: "slither",
		sound:      "hsss",
	}
}

func handleErrorIfAny(err error, terminate bool) {
	if err != nil {
		if terminate {
			log.Fatalln(err)
		} else {
			log.Println(err)
		}
	}
}

func (cow *Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow *Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow *Cow) Speak() {
	fmt.Println(cow.sound)
}

func (bird *Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird *Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird *Bird) Speak() {
	fmt.Println(bird.sound)
}

func (snake *Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake *Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake *Snake) Speak() {
	fmt.Println(snake.sound)
}

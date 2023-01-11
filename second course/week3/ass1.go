package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		request := scanner.Text()
		parts := strings.Split(request, " ")
		animal, info := parts[0], parts[1]

		a := animals[animal]

		switch info {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		}
	}
}

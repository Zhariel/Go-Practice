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

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if len(parts) < 2 {
			continue
		}

		animal, action := parts[0], parts[1]
		a, exists := animals[animal]

		if !exists {
			fmt.Println("Invalid animal")
			continue
		}

		switch action {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			fmt.Println("Invalid action")
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func insertAnimal(animals map[string]Animal, parts []string) {
	if len(parts) < 2 {
		return
	}

	animal, name := parts[0], parts[1]

	switch animal {
	case "cow":
		animals[name] = Cow{name}
	case "bird":
		animals[name] = Bird{name}
	case "snake":
		animals[name] = Snake{name}
	default:
		fmt.Println("Unrecognized animal")
		return
	}
	fmt.Println("Created it !")
}

func requestAnimal(animals map[string]Animal, parts []string) {
	if len(parts) < 2 {
		fmt.Println(parts)
		return
	}
	name, action := parts[0], parts[1]

	switch action {
	case "eat":
		animals[name].Eat()
	case "move":
		animals[name].Move()
	case "speak":
		animals[name].Speak()
	default:
		fmt.Println("Invalid action")
	}
}

func main() {
	animals := make(map[string]Animal)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		//command examples :
		//newanimal cow daisy
		//daisy move
		if parts[0] == "newanimal" {
			insertAnimal(animals, parts[1:])
		} else {
			requestAnimal(animals, parts)
		}
	}
}

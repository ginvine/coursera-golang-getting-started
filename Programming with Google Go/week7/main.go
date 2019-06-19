package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, sound string
}

func NewAnimal(food, locomotion, sound string) Animal {
	return Animal{
		food:       food,
		locomotion: locomotion,
		sound:      sound,
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func getAnimals() map[string]Animal {
	cow := NewAnimal("grass", "walk", "moo")
	bird := NewAnimal("worms", "fly", "peep")
	snake := NewAnimal("mice", "slither", "hsss")
	return map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
}

func main() {
	animals := getAnimals()
	for {
		fmt.Print("> ")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		line := s.Text()

		inputs := strings.Split(line, " ")
		if len(inputs) != 2 {
			fmt.Println("please enter exactly 2 words")
			continue
		}

		animal, ok := animals[inputs[0]]
		if !ok {
			fmt.Printf("can't find %s animal\n", inputs[0])
			continue
		}

		switch inputs[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("type eat, speak or move")
		}
	}
}

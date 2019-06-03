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
		r := bufio.NewReader(os.Stdin)
		b, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("incorrect input: %v\n", err)
		}
		inputs := strings.Split(string(b), " ")
		if len(inputs) != 2 {
			fmt.Println("please enter exactly 2 words")
		}
		animal, ok := animals[inputs[0]]
		if !ok {
			fmt.Printf("can't find %s animal\n", inputs[0])
		}
		switch inputs[1] {
		case "eat\n":
			animal.Eat()
		case "move\n":
			animal.Move()
		case "speak\n":
			animal.Speak()
		default:
			fmt.Println("unrecognized information request")
		}
	}
}

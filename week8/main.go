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

type Animals map[string]Animal

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func addNewAnimal(animals Animals, name, t string) error {
	switch t {
	case "cow":
		animals[name] = Cow{}
	case "bird":
		animals[name] = Bird{}
	case "snake":
		animals[name] = Snake{}
	default:
		return fmt.Errorf("incorrect animal type: %s", t)
	}
	fmt.Println("Created it!")
	return nil
}

func queryAnimal(animals Animals, name, info string) error {
	animal, ok := animals[name]
	if !ok {
		return fmt.Errorf("can't find animal with name %s", name)
	}

	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return fmt.Errorf("unknown information: %s", info)
	}
	return nil
}

func main() {
	animals := Animals{}

	for {
		fmt.Print("> ")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		line := s.Text()

		inputs := strings.Split(line, " ")
		if len(inputs) != 3 {
			fmt.Println("please enter exactly 3 words")
			continue
		}

		switch inputs[0] {
		case "newanimal":
			if err := addNewAnimal(animals, inputs[1], inputs[2]); err != nil {
				fmt.Println(err)
			}
		case "query":
			if err := queryAnimal(animals, inputs[1], inputs[2]); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Printf("unknown command: %s", inputs[0])
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname, lname string
}

func main() {

	fmt.Println("please enter a file name")
	var name string
	fmt.Scan(&name)

	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("Can't open file: %v", err)
	}
	defer file.Close()

	ppl := make([]person, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		ppl = append(ppl, person{split[0], split[1]})
	}

	for _, p := range ppl {
		fmt.Printf("%s %s\n", p.fname, p.lname)
	}
}

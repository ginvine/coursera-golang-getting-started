package main

import (
	"fmt"
	"sort"
	"strconv"
)

const exit = "X"

func main() {
	sli := make([]int, 3)

	for {
		var s string
		fmt.Printf("Enter an integer to be added to an array or \"%s\" for exit\n", exit)
		fmt.Scanf("%s", &s)

		if s == exit {
			return
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("incorrect input")
		}
		sli = append(sli, i)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}

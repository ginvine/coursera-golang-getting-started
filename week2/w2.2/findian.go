package main

import (
	"fmt"
	"strings"
)

const (
	start  = "i"
	end    = "n"
	middle = 'a'
)

func main() {
	var starts, ends, contains bool
	var s string

	fmt.Println("Please enter a string and hit ENTER")
	fmt.Scan(&s)
	s = strings.ToLower(s)

	starts = strings.HasPrefix(s, start)
	ends = strings.HasSuffix(s, end)
	contains = strings.ContainsRune(s, middle)

	if starts && ends && contains {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}

package main

import "fmt"

func main() {
	var input float64
	fmt.Println("Please enter floating point number and press ENTER")
	fmt.Scan(&input)
	fmt.Printf("Truncated version of the entered number is:\n%d\n", int(input))
}

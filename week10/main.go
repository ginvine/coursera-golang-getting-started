package main

import (
	"fmt"
)

func main() {
	x := 0
	go func() {
		//time.Sleep(time.Second)
		fmt.Println(x)
	}()
	go func() { x++ }()
	// time.Sleep(2 * time.Second)
	fmt.Println("done")
}

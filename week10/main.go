package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		//time.Sleep(time.Second)
		fmt.Println(x)
		wg.Done()
	}()
	go func() {
		x++
		wg.Done()
	}()
	// time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("done")
}

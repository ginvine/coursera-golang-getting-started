package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortArr(arr []int, wg *sync.WaitGroup) {
	fmt.Println(arr)
	sort.Ints(arr)
	wg.Done()
}

func main() {
	fmt.Println("vim-go")
	arr := []int{1, 3, 5, 6, 2, 1, 0, -4}
	fmt.Println(arr)
	wg := sync.WaitGroup{}
	wg.Add(4)
	a := 0
	inc := len(arr) / 4
	b := inc
	for i := 0; i < 3; i++ {
		go sortArr(arr[a:b], &wg)
		a = b
		b += inc
	}
	go sortArr(arr[a:], &wg)
	wg.Wait()
	sort.Ints(arr)
	fmt.Println(arr)
}

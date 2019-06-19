package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortArr(arr []int, wg *sync.WaitGroup) {
	fmt.Println(arr)
	sort.Ints(arr)
	wg.Done()
}

func main() {
	fmt.Println("enter numbers separated by space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := strings.Fields(scanner.Text())

	arr := make([]int, 0)
	for _, s := range str {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("can't convers %s to int: %v\n", s, err)
		}
		arr = append(arr, v)
	}

	fmt.Println("Sorting", arr)
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func Swap(arr []int, i int) {
	tmp := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = tmp
}

func main() {
	fmt.Println("enter numbers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := strings.Fields(scanner.Text())
	if len(str) > 10 {
		str = str[0:10]
	}

	arr := make([]int, 0)
	for _, s := range str {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("can't convers %s to int: %v\n", s, err)
		}
		arr = append(arr, v)
	}

	BubbleSort(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

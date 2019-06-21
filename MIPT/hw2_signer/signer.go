package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type job func(in, out chan interface{})

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	for _, j := range jobs {
		out := make(chan interface{})
		go func(j job, in, out chan interface{}) {
			j(in, out)
			close(out)
		}(j, in, out)
		in = out
	}
	for range in {
	}
}

func main() {
	jobs := []job{
		job(func(in, out chan interface{}) {
			fmt.Println("sending to chan")
			out <- 1
			fmt.Println("sent to chan")
		}),
		job(func(in, out chan interface{}) {
			fmt.Println("start reading from input chan")
			for data := range in {
				fmt.Println(data)
			}
			fmt.Println("stop reading from input chan")
		}),
	}
	ExecutePipeline(jobs...)
	//fmt.Scanln()
}

func SingleHash(in, out chan interface{}) {
	for input := range in {
		data := strconv.Itoa(input.(int))
		out <- DataSignerCrc32(data) + "~" + DataSignerCrc32(DataSignerMd5(data))
	}

}

func MultiHash(in, out chan interface{}) {
	for input := range in {
		data := input.(string)
		result := ""
		for i := 0; i < 6; i++ {
			result = result + DataSignerCrc32(strconv.Itoa(i)+data)
		}
	}
}

func CombineResults(in, out chan interface{}) {
	result := []string{}
	for input := range in {
		data := input.(string)
		result = append(result, data)
	}
	sort.Strings(result)
	out <- strings.Join(result, "_")
}

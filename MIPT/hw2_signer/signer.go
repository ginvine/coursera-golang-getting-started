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
	out := make(chan interface{})
	for _, j := range jobs {
		go func() {
			j(in, out)
			close(out)
		}()
		in = out
		out = make(chan interface{})
	}
}

func main() {
	jobs := []job{
		func(in, out chan interface{}) {
			out <- 1
		},
		func(in, out chan interface{}) {
			data := <-in
			fmt.Println(data)
		},
	}
	ExecutePipeline(jobs...)
	fmt.Scanln()
}

func SingleHash(in, out chan interface{}) {
	for input := range in {
		data := input.(string)
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

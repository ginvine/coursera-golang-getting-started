package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

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

func signCrc32(data string, out chan string) {
	out <- DataSignerCrc32(data)
}

func signMd5(data string, out chan string, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	out <- DataSignerMd5(data)
}

func SingleHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)
	for input := range in {
		data := strconv.Itoa(input.(int))
		wg.Add(1)
		go func(data string, out chan interface{}, wg *sync.WaitGroup) {
			var r1, r2 string
			ch1 := make(chan string)
			ch2 := make(chan string)
			go signCrc32(data, ch1)
			go signMd5(data, ch2, &mu)
			for i := 0; i < 2; i++ {
				select {
				case r1 = <-ch1:
				case r2 = <-ch2:
					go signCrc32(r2, ch2)
				}
			}
			r2 = <-ch2
			out <- r1 + "~" + r2
			wg.Done()
		}(data, out, wg)
	}
	wg.Wait()
}

type res struct {
	i int
	s string
}

func MultiHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)
	for input := range in {
		data := input.(string)
		wg.Add(1)
		go func(data string, out chan interface{}, wg *sync.WaitGroup) {
			ch := make(chan res)
			for i := 0; i < 6; i++ {
				go func(i int, data string, ch chan res) {
					r := DataSignerCrc32(strconv.Itoa(i) + data)
					ch <- res{i, r}
				}(i, data, ch)
			}
			s := make([]string, 6)
			for i := 0; i < 6; i++ {
				r := <-ch
				s[r.i] = r.s
			}
			out <- strings.Join(s, "")
			wg.Done()
		}(data, out, wg)
	}
	wg.Wait()
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

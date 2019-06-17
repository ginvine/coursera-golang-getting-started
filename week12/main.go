package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number                 int
	request, approve, done chan int
	leftCS, rightCS        *ChopS
	wg                     *sync.WaitGroup
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.requestPermission()
		<-p.approve

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		p.done <- p.number
	}

	fmt.Printf("finishing eating %d\n", p.number+1)

	p.wg.Done()
}

func (p Philo) requestPermission() {
	p.request <- p.number
}

func host(philos []*Philo, request, done chan int) {
	eating := 0
	for {
		if eating > 1 {
			<-done
			eating--
		}

		select {
		case <-done:
			eating--
		case n := <-request:
			eating++
			philos[n].approve <- 1
		}
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	request, done := make(chan int), make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(5)

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			number:  i,
			request: request,
			approve: make(chan int),
			done:    done,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
			wg:      wg,
		}
	}

	go host(philos, request, done)
	for i := range philos {
		go philos[i].eat()
	}

	wg.Wait()
}

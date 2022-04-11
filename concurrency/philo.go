package main

import (
	"fmt"
	"sync"
)

type chopS struct {
	sync.Mutex
}

type Philosopher struct {
	id              int
	leftCS, rightCS *chopS
}

func eating(wg *sync.WaitGroup, channelPhilosophers <-chan Philosopher) {
	for x := range channelPhilosophers {
		x.leftCS.Lock()
		x.rightCS.Lock()
		fmt.Println("starting to eat", x.id+1)
		x.rightCS.Unlock()
		x.leftCS.Unlock()
		fmt.Println("finishing eating", x.id+1)
		wg.Done()
	}

}

func host(philosophers []*Philosopher, wg *sync.WaitGroup) {
	channelPhilosophers := make(chan Philosopher, 10)
	for i := 0; i < 5; i++ {
		channelPhilosophers <- *philosophers[i]
		go eating(wg, channelPhilosophers)
	}

}

func main() {
	var wg sync.WaitGroup
	CSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopS)
	}
	philosophers := make([]*Philosopher, 5) //storage Philosopher
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:      i,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
		}
	}
	wg.Add(15)
	for i := 0; i < 3; i++ {
		host(philosophers, &wg)
	}
	wg.Wait()

}

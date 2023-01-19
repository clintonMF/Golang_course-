package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type chopS struct{ sync.Mutex }

type philo struct {
	leftCs, rightCs *chopS
	num             int
}

func (p philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Println("starting to eat", p.num)

		fmt.Printf("finishing eating %v\n\n", p.num)

		p.leftCs.Unlock()
		p.rightCs.Unlock()
	}

	wg.Done()
}

func main() {
	CSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopS)
	}

	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		if i < 5 {
			philos[i] = &philo{CSticks[i], CSticks[(i+1)%5], i}
		} else {
			philos[i] = &philo{CSticks[(i+1)%5], CSticks[i], i}

		}
	}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Philosopher struct {
	leftChopstick, rightChopstick *ChopStick
	nbTimesEaten, id              int
}

func (p *Philosopher) eat(wg *sync.WaitGroup, host *chan int) {
	for p.nbTimesEaten < 3 {

		var first, second = pickRandomChopstickOrdering(p)

		(*host) <- p.id

		first.Lock()
		second.Lock()
		fmt.Printf("Starting to eat %v\n", p.id)

		p.nbTimesEaten += 1
		// fmt.Printf("Philosopher %v has eaten %v times\n", p.id, p.nbTimesEaten)

		fmt.Printf("finishing eating %v\n", p.id)
		first.Unlock()
		second.Unlock()

		<-(*host)
	}

	wg.Done()
}

func pickRandomChopstickOrdering(p *Philosopher) (*ChopStick, *ChopStick) {
	if rand.Float64() > 0.5 {
		return p.leftChopstick, p.rightChopstick
	} else {
		return p.rightChopstick, p.leftChopstick
	}
}

type ChopStick struct{ sync.Mutex }

func main() {
	var wg sync.WaitGroup

	host := make(chan int, 2)

	chopsticks := make([]*ChopStick, 5)
	philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = &ChopStick{}
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			nbTimesEaten:   0,
			id:             i + 1,
		}
	}

	wg.Add(5)
	for _, philosopher := range philosophers {
		go philosopher.eat(&wg, &host)
	}

	wg.Wait()

	fmt.Println("\nVerification: ")
	for _, philosopher := range philosophers {
		fmt.Printf("Philosopher %v has eaten %v times\n", philosopher.id, philosopher.nbTimesEaten)
	}
}

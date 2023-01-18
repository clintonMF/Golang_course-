/*
- Sharing variables concurrently can cause problems
- 2 go routines writing to a shared variable can interfere
with one another.

Concurrency-safe
- A function is said to be concurrency safe if it can be invoked
concurrently without interfering with other goroutines
*/

/*
* CORRECT SHARING
- don't let two go routines write to the same variable at the same
time
- Need to restrict possible interleaving
- Access to shared variable cannot be interleaved
* MUTUAL EXCLUSION
- Code segment in different go routines which cannot execute concurrently
- writing to shared variable should be mutually exclusive

the above principles are implemented in golang with the sync.Mutex
- a mutex ensures mutual exclusion
- it uses a binary semaphore
- Lock() is used to show that a go routine is currently accessing the
shared variable and will block other go routine from accessing it until the
first go routine is done.

example is shown below
*/

package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex
var i int = 0
var wg sync.WaitGroup

func inc() {
	mut.Lock()
	i = i + 1
	fmt.Println(i)
	mut.Unlock()
	wg.Done()
}

func pri() {
	fmt.Println("here")
	wg.Done()
}

func main() {
	wg.Add(3)
	go inc()
	go inc()
	go pri()

	wg.Wait()
	fmt.Println(i)
}

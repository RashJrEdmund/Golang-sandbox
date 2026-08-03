/*
MUTUAL EXCLUSION:
	A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	A Mutex must not be copied after first use.
	Mutexes allow us to lock access to data.
	-	This ensures that we can control which goroutines can access certain data at which time.
*/

/*
	SYNC.MUTEX: https://pkg.go.dev/sync#Mutex
	Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type
	and its two methods:
	-	.Lock()
	-	.Unlock()

	It's advisable to structure the protected code within a function so that the defer statement is used to unlock the mutex.

		func() {
			mu.Lock()
			defer mu.Unlock()
			// protected code
		}()
*/

/*
	SYNC.RWMUTEX: https://pkg.go.dev/sync#RWMutex
	Go's standard library provides a built-in implementation of a read-write mutex with the sync.RWMutex type
	and its four methods:
	-	.Lock()
	-	.Unlock()
	-	.RLock()
	-	.RUnlock()

		RLock & RUnlock are used to lock and unlock the mutex for concurrent read operations.
		Which improves performance in read-intensive processes. By allowing multiple goroutines to safely read
		from the map simultaneously. as many RLock() calls can occur at the same time.
		However, only one goroutine can hold a Lock(), and during this time, all RLock() operations are blocked.

	so instead of using sync.Mutex's Lock and Unlock for both read and write operations,
		we can use sync.RWMutex's RLock and RUnlock for read operations, and Lock and Unlock for write operations.

		func() {
			mu.RLock()
			defer mu.RUnlock()
			// protected code that reads from the map
		}()

		func() {
			mu.Lock()
			defer mu.Unlock()
			// protected code that writes to the map
		}()
*/

/*
MAPS ARE NOT THREAD-SAFE:
	Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map,
		and at least one of them is writing to the map, you must lock your maps with a mutex.

		mu := sync.Mutex{}
		m := make(map[string]int)

		go func() {
			mu.Lock()
			defer mu.Unlock()
			m["key"] = 1

Read Thread Safety Here:
	https://en.wikipedia.org/wiki/Thread_safety

Checkout the following link for more information on atomic maps:
	https://go.dev/doc/faq#atomic_maps

Read Mutual Exclusions Here:
	https://en.wikipedia.org/wiki/Mutual_exclusion
*/

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Mutexes")
// }

package main

import (
	"fmt"
	"sync"
)

type IntMap = map[int]int

func main() {
	m := IntMap{}

	mu := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	readLoops := []func(IntMap, *sync.RWMutex, *sync.WaitGroup, int){ // an array of readLoop functions
		readLoop,
		readLoop,
		readLoop,
		readLoop,
	}

	wg.Add(1)
	go writeLoop(m, mu, wg)

	for tag, readLoop := range readLoops {
		wg.Add(1)
		go readLoop(m, mu, wg, tag)
	}

	wg.Wait()
}

func writeLoop(m IntMap, mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer mu.Unlock()
	defer wg.Done()
	/*
		defer is executed in LIFO order. so wg.Done() will be executed before the mu.Unlock()

		if we were to write:
			defer wg.Done()
			defer mu.Unlock()

		then mu.Unlock() would be executed before the wg.Done().
		Which would cause a race condition:
			There wil be a tiny gap of time where the data is unlocked but the main goroutine tracking the WaitGroup still thinks this worker is busy.
			During this gap, another goroutine could try to grab the lock and modify the data, causing a race condition.

		Letting the WaitGroup know a task is done before releasing its internal locks is a race-condition anti-pattern
			Always release resource locks as the very last step.
	*/
	/*
		If we tried using an anonymous func in defer to Done and unlock,
			defer func(){
				wg.Done()   // Line 1 -> Runs 1st
				mu.Unlock() // Line 2 -> Runs 2nd
			}()
		This keeps the sequence correct but has a hidden gotcha:
			If wg.Done() panics inside the closure, the execution of that specific anonymous function hits an unrecovered panic state and halts.
			The next line (mu.Unlock()) never executes. The mutex stays locked forever, deadlocking the entire program.
	*/

	mu.Lock() // lock the mutex. writing to the map is a critical section.
	for i := 1; i <= 100; i++ {
		m[i] = i
		fmt.Println("writing", i)
	}

	fmt.Println("Done writing")
}

func readLoop(m IntMap, mu *sync.RWMutex, wg *sync.WaitGroup, tag int) {
	defer mu.RUnlock()
	defer wg.Done()

	fmt.Printf("\nReading In go routine %d: %d items\n", tag, len(m))

	mu.RLock()
	for k, v := range m {
		fmt.Println(tag, ":", k, "-", v)
	}
}

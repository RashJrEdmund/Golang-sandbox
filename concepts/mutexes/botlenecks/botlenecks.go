/*
Mutexes synchronize access to shared resources, but they can also create bottlenecks if not used correctly.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(2 * time.Second)
	return data * 2
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()

	lock.Lock()

	processedData := process(data)
	*result = append(*result, processedData) // critical section

	lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	start := time.Now()

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println("Time taken: ", time.Since(start))
	fmt.Println("Result: ", result)
}

/*
Output:
Time taken:  10.000211971s
Result:  [2 4 6 8 10]

This is a bottleneck because the mutex is locking and unlocking the data for each iteration.

so it's basically synchronous processing of the data.

To fix this, we can use a confinement pattern to limit the scope of the mutex to the data we are processing.

Watch this video for solution:
	-> start here to better understand: https://youtu.be/Bk1c30avsuU?t=955
	-> or skip directly to confinement pattern: https://www.youtube.com/watch?v=Bk1c30avsuU&t=1073s
*/

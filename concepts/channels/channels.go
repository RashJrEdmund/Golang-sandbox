/*
X POST:
https://x.com/orashus/status/2083577936368566364
*/
/*
Channels are a typed, thread-safe queue.
Channels allow different goroutines to communicate with each other.

Channels are a way to communicate between goroutines.
They are a way to synchronize goroutines.

The <- operator is called the channel operator. Data flows in the direction of the arrow.
	This operation will block the go routine/function that is sending the value until another goroutine is ready to receive the value or
	there is a value in the channel to be read

Channels are First-In-First-Out (FIFO), meaning values are received in the same order they are sent.

Like maps and slices, channels are reference types

Example:
func send(ch chan int) {
	ch <- 99
}

func main() {
	ch := make(chan int)
	go send(ch)
	fmt.Println(<-ch) // 99
}
*/

/*
	FOR UNBUFFERED CHANNELS:
 	Readers should always be guaranteed before sending except the sender is a go routine ?

  1. Sender first, no reader yet → DEADLOCK

		ch := make(chan int)
		ch <- 42  // main blocks here
		go func(){ <-ch }() // never runs

	2. Reader first, then sender → WORKS

		ch := make(chan int)
		go func(){ <-ch }() // reader waiting
		ch <- 42  // main hands it off

	3. Sender is in a goroutine → ALSO WORKS

		ch := make(chan int)
		go func(){ ch <- 42 }() // sender goroutine blocks, but doesn't stop main
		v := <-ch // main reads, unblocks the sender
*/

/*
	FOR BUFFERED CHANNELS:
	With buffered channels it's different.
	Created with make(chan T, N) gives you N "free" sends before you need a reader.

		ch := make(chan int, 3)
		ch <- 1 // ok, no reader needed yet
		ch <- 2 // ok
		ch <- 3 // ok
		ch <- 4 // now blocks until someone reads
*/

package main

import (
	"fmt"
	// "sync"
)

func main() {
	ch := make(chan int)
	fmt.Println("In main")

	go func() {
		fmt.Println("In goroutine")
		v := <-ch
		fmt.Println("Got", v)
	}()

	fmt.Println("Before sending")
	ch <- 1
	fmt.Println("Sent And Done")
}

// func main() {
// 	ch := make(chan int)
// 	wg := sync.WaitGroup{}

// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()

// 		fmt.Println("In A")
// 		ch <- 42
// 		fmt.Println("A finished")
// 	}() // Goroutine A: tries to send, blocks because no reader yet

// 	go func() {
// 		defer wg.Done()

// 		fmt.Println("In B")
// 		v := <-ch
// 		fmt.Println("B got", v)
// 		fmt.Println("B finished")
// 	}()

// 	wg.Wait()
// }

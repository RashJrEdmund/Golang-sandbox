/*
X POST:
https://x.com/orashus/status/2083577936368566364
*/
/*
CHANNEL AXIOMS:
https://dave.cheney.net/2014/03/19/channel-axioms
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

	you can use the len() function to get the number of items in the channel
		and the cap() function to get the capacity of the channel

		channel := make(chan int, 3)
		channel <- 1
		len(channel) // 1
		cap(channel) // 3
*/

/*
CLOSING CHANNELS:
Channels can be explicitly closed by a sender:

	ch := make(chan int)

	// do some stuff with the channel

	close(ch)

CHECKING IF A CHANNEL IS CLOSED:
Similar to the ok value when accessing data in a map,
receivers can check the ok value when receiving from a channel to test if a channel was closed.

	v, ok := <-ch // channel is closed if ok is false

Sending on a closed channel will cause a panic.
Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused.
You should close channels to indicate explicitly to a receiver that nothing else is going to come across.
*/

/*
SELECT:
Sometimes we have a single goroutine listening to multiple channels
and want to process data in the order it comes through each channel.

A select statement is used to listen to multiple channels at the same time.
It is similar to a switch statement but for channels.

	select {
	case i, ok := <-chInts:
		if ok {
			fmt.Println(i)
		}
	case s, ok := <-chStrings:
		if ok {
			fmt.Println(s)
		}
	default:
		// The default case in a select statement executes immediately if no other channel has a value ready.
		// A default case stops the select statement from blocking
	}

The first channel with a value ready to be received will fire and its body will execute
*/

/*
READ ONLY CHANNELS:
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:
	func main() {
		ch := make(chan int)
		readCh(ch)
	}

	func readCh(ch <-chan int) {
		// ch can only be read from
		// in this function
	}

WRITE-ONLY CHANNELS:
The same goes for write-only channels, but the arrow's position moves.
	func writeCh(ch chan<- int) {
		// ch can only be written to
		// in this function
	}
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

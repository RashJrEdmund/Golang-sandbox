/*
Go routines are lightweight threads of execution.

They are managed by the Go runtime and are not associated with any specific OS thread.

They are used to implement concurrency in Go.

They are created using the go keyword.

They are started using the go keyword.

They are stopped using the go keyword.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	tableNumber int
	order       string
	prepTime    time.Duration
}

func processOrder(order Order) {
	fmt.Printf("Processing Order \"%s\" for table %d\n", order.order, order.tableNumber)
	fmt.Println("Estimated time:", order.prepTime.Seconds(), "seconds")

	time.Sleep(order.prepTime)

	fmt.Printf("Order for table %d is ready\n", order.tableNumber)
}

func main() {
	orders := []Order{
		{tableNumber: 1, prepTime: 2 * time.Second, order: "Salad"},
		{tableNumber: 2, prepTime: 3 * time.Second, order: "Pizza"},
		{tableNumber: 3, prepTime: 1 * time.Second, order: "Ice Cream"},
		{tableNumber: 4, prepTime: 5 * time.Second, order: "Steak"},
		{tableNumber: 5, prepTime: 4 * time.Second, order: "Fried Chicken"},
	}

	wg := sync.WaitGroup{}
	start := time.Now()

	// wg.Add(len(orders))

	for _, order := range orders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processOrder(order)
		}()
	}

	wg.Wait()

	fmt.Println("\nTotal Time", time.Since(start))
}

/*
go routines and wait groups is kinda like awaiting a Promise.All([]) in JavaScript.
*/

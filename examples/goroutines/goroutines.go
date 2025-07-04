package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// say prints the given string multiple times with a delay.
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Goroutines demonstrates the use of goroutines for concurrent execution and basic channel communication.
func Goroutines() {
	fmt.Println("\n--- Goroutines Example ---")

	var wg sync.WaitGroup

	// Example 1: Basic Goroutines with WaitGroup
	wg.Add(1)
	go say("world", &wg) // Run say("world") in a new goroutine

	wg.Add(1)
	go say("hello", &wg) // Run say("hello") in a new goroutine

	// Wait for the first set of goroutines to finish
	wg.Wait()

	// Example 2: Anonymous Goroutine with WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Anonymous goroutine says hi!")
	}()

	// Wait for the anonymous goroutine to finish
	wg.Wait()

	// Example 3: Using channels for communication (brief introduction)
	messages := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		messages <- "ping" // Send a message to the channel
	}()

	msg := <-messages // Receive a message from the channel
	fmt.Println(msg)

	// Wait for the channel goroutine to finish
	wg.Wait()
}

package goroutines

import (
	"fmt"
	"time"
)

// say prints the given string multiple times with a delay.
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Goroutines demonstrates the use of goroutines for concurrent execution and basic channel communication.
func Goroutines() {
	go say("world") // Run say("world") in a new goroutine
	say("hello")    // Run say("hello") in the main goroutine

	// We need to wait a bit for the goroutines to finish,
	// otherwise the main goroutine might exit before they complete.
	time.Sleep(time.Second)

	// Another example with anonymous function
	go func() {
		fmt.Println("Anonymous goroutine says hi!")
	}()
	time.Sleep(100 * time.Millisecond)

	// Using channels for communication (brief introduction)
	messages := make(chan string)

	go func() {
		messages <- "ping" // Send a message to the channel
	}()

	msg := <-messages // Receive a message from the channel
	fmt.Println(msg)
}

package main

import (
	"fmt"
	"sync"
)

func printSomething(something string, waitgroup *sync.WaitGroup) {
	// Notify the waitgroup that this goroutine is done when the function exits.
	defer waitgroup.Done()
	fmt.Println(something)
}

// main is the entry point of the application.
// It's a go routine that runs when the program starts.
func main() {
	
	var waitgroup sync.WaitGroup

	words := []string{"alpha", "beta", "gamma", "delta", "pi", "theta", "lambda", "omega"}

	// Set the number of goroutines to wait for.
	waitgroup.Add(len(words))

	for index, word := range words {
		go printSomething(fmt.Sprintf("Word %d: %s", index+1, word), &waitgroup)
	}

	waitgroup.Wait()

	// Start a new goroutine to print the first message.
	// go printSomething("This is the first thing to be printed.", &waitgroup)

	// time.Sleep(1 * time.Second) // Add a sleep to allow the goroutine to complete before main exits.	
								// It's generally not a good practice to use sleep for synchronization,
								// but it's used here for simplicity in this example.

	// Start a new goroutine to print the second message.
	waitgroup.Add(1)

	printSomething("This is the second thing to be printed.", &waitgroup)
}
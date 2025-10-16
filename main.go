package main

import (
	"fmt"
	"time"
)

func printSomething(something string) {
	fmt.Println(something)
}

// main is the entry point of the application.
// It's a go routine that runs when the program starts.
func main() {
	
	go printSomething("This is the first thing to be printed.")

	time.Sleep(1 * time.Second) // Add a sleep to allow the goroutine to complete before main exits.	
								// It's generally not a good practice to use sleep for synchronization,
								// but it's used here for simplicity in this example.
	printSomething("This is the second thing to be printed.")
}
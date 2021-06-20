package main

import (
	"bufio"
	"fmt"
	"os"
)

// main function
func main() {
	// string for storing input
	var input string

	// Slice for storing the history of the world.
	var history []string

	fmt.Println("Hello, World!")
	fmt.Println("I am Echo! Please tell me something to say by typing it in, and pressing enter!")

	// Stop when we see "bye"
	for input != "bye" {

		fmt.Print("You: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = scanner.Text()
			break
		}
		// Add the new input to the history
		history = append(history, input)
		fmt.Print("Echo: " + input + "\n")

		// Check if the input was the string 'history'
		if input == "history" {
			// iterate (loop) through the history
			for k, v := range history {
				// print out the Key and the Value
				fmt.Println(k, v)
			}
		}
	}

}

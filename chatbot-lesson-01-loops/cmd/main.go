package main

import (
	"bufio"
	"fmt"
	"os"
)

// main function
func main() {
	var input string

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

		fmt.Print("Echo: " + input + "\n")
	}
}

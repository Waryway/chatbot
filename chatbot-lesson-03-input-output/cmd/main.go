package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// main function
func main() {
	f, err := os.OpenFile("history.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// check for the error.
	if err != nil { // a 'nil' error means no error occurred.
		fmt.Println(err) // print the error
		return           // stop processing the context.
	}

	defer func() {
		_ = f.Close()
	}()

	// string for storing input
	var input string

	// Slice for storing the history of the world.
	var history []string
	scanner := bufio.NewScanner(f) // use the buffer input output package to start a file scanner
	for scanner.Scan() {           // loop through the file scanner
		history = append(history, scanner.Text()) // add the scanned line to the lines slice.
	}

	fmt.Println("Hello, World!")
	fmt.Println("I am Echo! Please tell me something to say by typing it in, and pressing enter!")
	_, _ = f.WriteString("Hello, World! \n")
	_, _ = f.WriteString("I am Echo! Please tell me something to say by typing it in, and pressing enter! \n")
	// Stop when we see "bye"
	for input != "bye" {

		fmt.Print("You: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = scanner.Text()
			break
		}
		_, _ = f.WriteString("You: " + input + "\n")
		// Add the new input to the history
		history = append(history, input)
		fmt.Print("Echo: " + input + "\n")
		_, _ = f.WriteString("Echo: " + input + "\n")
		// Check if the input was the string 'history'
		if input == "history" {
			// We won't add the history to the history, instead we can add how many lines of history exist.
			_, _ = f.WriteString("<Truncated> " + strconv.Itoa(len(history)) + " Lines of history repetition.")
			// iterate (loop) through the history
			for k, v := range history {
				// print out the Key and the Value
				fmt.Println(k, v)
			}
		}
	}
}

package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type EchoBot struct {
	greeting   string
	userPrompt string
	history    History
}

func (e EchoBot) Initialize(history History) EchoBot {
	e.greeting = "Hello, World! \n I am Echo! Please tell me something to say by typing it in, and pressing enter!"
	e.userPrompt = ">"
	e.history = history
	return e
}

func (e EchoBot) GreetUser() EchoBot {
	e.history.Record(e.greeting, true)
	return e
}

func (e EchoBot) ChatLoop() EchoBot {
	// string for storing input
	var input string

	// Stop when we see "bye"
	for input != "bye" {
		e.prompt()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = scanner.Text()
			break
		}

		// Add the new input to the history
		e.history.Record(e.userPrompt+input, false)
		e.history.Record("Echo: "+input, true)
		// Check if the input was the string 'history'
		if input == "history" {
			// We won't add the history to the history, instead we can add how many lines of history exist.
			e.history.Record("<Truncated> "+strconv.Itoa(e.history.Length())+" Lines of history repetition.", false)
			e.history.Print()
		}
	}
	return e
}

func (e EchoBot) prompt() {
	fmt.Print(e.userPrompt)
}

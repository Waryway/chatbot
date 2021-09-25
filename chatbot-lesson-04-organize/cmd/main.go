package main

import (
	"fmt"
	"github.com/waryway/chatbot/chatbot-lesson-04-organize/internal/handlers"
)

// main function
func main() {
	var history handlers.History
	err := history.Initialize("history.txt").LoadHistory()

	// check for the error.
	if err != nil { // a 'nil' error means no error occurred.
		fmt.Println(err) // print the error
		return           // stop processing the context.
	}

	var echo handlers.EchoBot
	echo.Initialize(history).GreetUser().ChatLoop()
}

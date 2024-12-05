package main

import (
	"github.nubank.com/capital-gains/console"
	"github.nubank.com/capital-gains/logic"
)

// main is the entry point of the application.
func main() {
	// Create a new instance of the logic layer.
	l := logic.NewLogic()

	// Create a new console interface with the logic processor.
	c := console.NewConsole(l)

	// Run the console to start processing user input.
	c.Run()
}

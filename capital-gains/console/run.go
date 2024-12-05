package console

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.nubank.com/capital-gains/logic"
	"os"
)

// Run handles the main console interaction, processing user input line by line.
func (c *console) Run() {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read from standard input.

	// Read input line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Exit the loop if the input line is empty
		if line == "" {
			break
		}

		// Parse the input JSON into a slice of operations
		var operations []logic.Operation
		err := json.Unmarshal([]byte(line), &operations)
		if err != nil {
			// Print error message to standard error if JSON parsing fails
			fmt.Fprintln(os.Stderr, "Error parsing input. Valid input must be a JSON array of operations:", err)
			continue
		}

		// Process the operations using the processor
		results, err := c.processor.Process(operations)
		if err != nil {
			// Print error message to standard error if processing fails
			fmt.Fprintln(os.Stderr, "Error processing input:", err)
			continue
		}

		// Marshal the results into JSON and print to standard output
		output, err := json.Marshal(results)
		if err != nil {
			// Print error message to standard error if JSON marshaling fails
			fmt.Fprintln(os.Stderr, "Error generating output JSON:", err)
			continue
		}

		fmt.Println(string(output)) // Print the results as JSON to standard output
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

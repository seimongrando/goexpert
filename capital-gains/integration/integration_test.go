package integration

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.nubank.com/capital-gains/console"
	"github.nubank.com/capital-gains/logic"
)

func Test_Integration(t *testing.T) {
	// Embedded input data
	inputData := `[{"operation":"buy", "unit-cost":10.00, "quantity":10000},{"operation":"sell", "unit-cost":15.00, "quantity":50},{"operation":"sell", "unit-cost":15.00, "quantity":50}]`

	// Expected output
	expectedOutput := `[{"tax":0.00},{"tax":0.00},{"tax":0.00}]`

	// Simulate stdin with the input data
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	stdinR, stdinW, _ := os.Pipe()
	os.Stdin = stdinR
	stdinW.WriteString(inputData + "\n\n")
	stdinW.Close()

	// Capture stdout
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	stdoutR, stdoutW, _ := os.Pipe()
	os.Stdout = stdoutW

	// Setup logic and console
	l := logic.NewLogic()
	c := console.NewConsole(l)

	// Run the console
	c.Run()

	// Read the output
	stdoutW.Close()
	var output bytes.Buffer
	output.ReadFrom(stdoutR)

	// Validate the output
	assert.JSONEq(t, expectedOutput, output.String(), "unexpected integration output")
}

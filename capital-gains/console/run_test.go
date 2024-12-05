package console

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"

	"github.nubank.com/capital-gains/logic"
)

func Test_Run(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		mockSetup      func(m *mocks)
		expectedOutput string
		expectedError  string
	}{
		{
			name:  "process valid input successfully",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity":100},{"operation":"sell", "unit-cost":50.00, "quantity":100}]`,
			mockSetup: func(m *mocks) {
				operations := logic.Operations{
					{Operation: "buy", UnitCost: 10.00, Quantity: 100},
					{Operation: "sell", UnitCost: 50.00, Quantity: 100},
				}
				results := []logic.TaxResult{
					{Tax: 0.00},
					{Tax: 8000.00},
				}
				m.processor.EXPECT().Process(operations).Return(results, nil)
			},
			expectedOutput: `[{"tax":0.00},{"tax":8000.00}]`,
		},
		{
			name:          "handle invalid JSON input",
			input:         `invalid json`,
			mockSetup:     func(m *mocks) {}, // No mock setup needed
			expectedError: "Error parsing input. Valid input must be a JSON array of operations:",
		},
		{
			name:  "handle processing error",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity":100}]`,
			mockSetup: func(m *mocks) {
				operations := logic.Operations{
					{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				}
				m.processor.EXPECT().Process(operations).Return(nil, errors.New("mock processing error"))
			},
			expectedError: "Error processing input: mock processing error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			console, mocks := createMocks(t)

			// Setup mocks
			if tt.mockSetup != nil {
				tt.mockSetup(mocks)
			}

			// Capture stdout and stderr using os.Pipe
			oldStdout := os.Stdout
			oldStderr := os.Stderr
			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			// Simulate input using os.Pipe
			oldStdin := os.Stdin
			stdinR, stdinW, _ := os.Pipe()
			os.Stdin = stdinR
			stdinW.WriteString(tt.input + "\n\n")
			stdinW.Close()

			// Run the console
			console.Run()

			// Close pipes and restore originals
			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr
			os.Stdin = oldStdin

			// Read outputs
			var stdoutOutput bytes.Buffer
			var stderrOutput bytes.Buffer
			stdoutOutput.ReadFrom(rOut)
			stderrOutput.ReadFrom(rErr)

			// Validate output or error
			if tt.expectedError != "" {
				assert.Contains(t, stderrOutput.String(), tt.expectedError)
			} else {
				assert.JSONEq(t, tt.expectedOutput, strings.TrimSpace(stdoutOutput.String()))
			}
		})
	}
}

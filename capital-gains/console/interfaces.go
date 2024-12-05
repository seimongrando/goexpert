package console

import "github.nubank.com/capital-gains/logic"

//go:generate mockgen -source=interfaces.go -destination=interfaces_test.go -package=console

// processor defines the interface for processing financial operations.
// This abstraction allows for easy mocking and testing of the console logic.
type processor interface {
	// Process handles a series of operations and returns the calculated tax results or an error.
	Process(operations logic.Operations) ([]logic.TaxResult, error)
}

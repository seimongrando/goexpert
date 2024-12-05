package console

// console represents the main interface for processing user input and displaying results.
type console struct {
	processor processor // Dependency responsible for processing operations.
}

// NewConsole creates and returns a new instance of console.
func NewConsole(processor processor) *console {
	return &console{processor: processor}
}

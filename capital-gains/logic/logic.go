package logic

// logic represents the core processing logic for financial operations.
type logic struct{}

// NewLogic creates and returns a new instance of logic.
// This is useful for dependency injection and future extensibility.
func NewLogic() *logic {
	return &logic{}
}

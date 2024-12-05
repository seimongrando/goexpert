package logic

import "fmt"

// OperationType defines the type of operation as a string (buy or sell).
type OperationType string

// Operations is a slice of Operation, representing multiple operations.
type Operations []Operation

// Constants for the types of operations and tax-free limit.
const (
	buyOperation  OperationType = "buy"    // Represents a buy operation.
	sellOperation OperationType = "sell"   // Represents a sell operation.
	totalTaxFree                = 20000.00 // Maximum value for a tax-free operation.
)

// Operation represents a financial operation with a type, unit cost, and quantity.
type Operation struct {
	Operation string  `json:"operation"` // Type of operation (buy or sell).
	UnitCost  float64 `json:"unit-cost"` // Cost per unit.
	Quantity  int     `json:"quantity"`  // Quantity involved in the operation.
}

// TaxResult represents the tax result of an operation.
type TaxResult struct {
	Tax float64 `json:"-"` // The calculated tax, hidden by default in JSON serialization.
}

// Total calculates the total value of the operation (unit cost * quantity).
func (o *Operation) Total() float64 {
	return o.UnitCost * float64(o.Quantity)
}

// TaxFree checks if the operation is exempt from tax based on its total value.
func (o *Operation) TaxFree() bool {
	return o.Total() <= totalTaxFree
}

// MarshalJSON customizes the JSON output for TaxResult, ensuring the tax is always formatted to two decimal places.
func (t TaxResult) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"tax":%.2f}`, t.Tax)), nil
}

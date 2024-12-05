package logic

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Process(t *testing.T) {
	tests := []struct {
		name       string
		operations []Operation
		expected   []TaxResult
	}{
		{
			name: "Case #1 - No Tax for Small Transactions",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
			},
		},
		{
			name: "Case #2 - Profit and Loss Deduction",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 10000.00},
				{Tax: 0.00},
			},
		},
		{
			name: "Case #3 - Loss Followed by Partial Profit",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 1000.00},
			},
		},
		{
			name: "Case #4 - No Profit or Loss with Weighted Average",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
			},
		},
		{
			name: "Case #5 - Mixed Transactions with Final Profit",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 10000.00},
			},
		},
		{
			name: "Case #6 - Complex Case with Loss Deduction",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 3000.00},
			},
		},
		{
			name: "Case #7 - Complex Case with New Buy and Multiple Deductions",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
				{Operation: "sell", UnitCost: 30.00, Quantity: 650},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 3000.00},
				{Tax: 0.00},
				{Tax: 0.00},
				{Tax: 3700.00},
				{Tax: 0.00},
			},
		},
		{
			name: "Case #8 - Large Profits with New Buy",
			operations: []Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			},
			expected: []TaxResult{
				{Tax: 0.00},
				{Tax: 80000.00},
				{Tax: 0.00},
				{Tax: 60000.00},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logic := NewLogic()
			results, err := logic.Process(tc.operations)
			require.NoError(t, err, "unexpected error: %v", err)

			// Convert results to JSON for easy comparison
			gotJSON, _ := json.Marshal(results)
			expectedJSON, _ := json.Marshal(tc.expected)
			require.Equal(t, string(expectedJSON), string(gotJSON), "expected %s, got %s", expectedJSON, gotJSON)
		})
	}
}

func Test_UpdateWeightedAverage(t *testing.T) {
	tests := []struct {
		name            string
		weightedAverage float64
		totalQuantity   int
		unitCost        float64
		quantity        int
		expectedAverage float64
		expectedTotal   int
	}{
		{
			name:            "Initial buy",
			weightedAverage: 0.0,
			totalQuantity:   0,
			unitCost:        10.0,
			quantity:        100,
			expectedAverage: 10.0,
			expectedTotal:   100,
		},
		{
			name:            "Additional buy",
			weightedAverage: 10.0,
			totalQuantity:   100,
			unitCost:        20.0,
			quantity:        50,
			expectedAverage: 13.33, // ((100*10) + (50*20)) / 150
			expectedTotal:   150,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			avg, total := updateWeightedAverage(tc.weightedAverage, tc.totalQuantity, tc.unitCost, tc.quantity)
			assert.InEpsilon(t, tc.expectedAverage, avg, 0.01, "unexpected weighted average")
			assert.Equal(t, tc.expectedTotal, total, "unexpected total quantity")
		})
	}
}

func Test_CalculateProfit(t *testing.T) {
	tests := []struct {
		name       string
		totalValue float64
		totalCost  float64
		expected   float64
	}{
		{
			name:       "Profit scenario",
			totalValue: 1000.0,
			totalCost:  800.0,
			expected:   200.0,
		},
		{
			name:       "No profit",
			totalValue: 800.0,
			totalCost:  800.0,
			expected:   0.0,
		},
		{
			name:       "Loss scenario",
			totalValue: 700.0,
			totalCost:  800.0,
			expected:   -100.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := calculateProfit(tc.totalValue, tc.totalCost)
			assert.Equal(t, tc.expected, result, "unexpected profit")
		})
	}
}

func Test_AccumulateLoss(t *testing.T) {
	tests := []struct {
		name      string
		totalLoss float64
		profit    float64
		expected  float64
	}{
		{
			name:      "No previous loss",
			totalLoss: 0.0,
			profit:    -100.0,
			expected:  100.0,
		},
		{
			name:      "Accumulate existing loss",
			totalLoss: 50.0,
			profit:    -150.0,
			expected:  200.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := accumulateLoss(tc.totalLoss, tc.profit)
			assert.Equal(t, tc.expected, result, "unexpected accumulated loss")
		})
	}
}

func Test_DeductLoss(t *testing.T) {
	tests := []struct {
		name           string
		profit         float64
		totalLoss      float64
		expectedProfit float64
		expectedLoss   float64
	}{
		{
			name:           "Profit greater than loss",
			profit:         100.0,
			totalLoss:      50.0,
			expectedProfit: 50.0,
			expectedLoss:   0.0,
		},
		{
			name:           "Profit equal to loss",
			profit:         100.0,
			totalLoss:      100.0,
			expectedProfit: 0.0,
			expectedLoss:   0.0,
		},
		{
			name:           "Profit less than loss",
			profit:         50.0,
			totalLoss:      100.0,
			expectedProfit: 0.0,
			expectedLoss:   50.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			remainingProfit, remainingLoss := deductLoss(tc.profit, tc.totalLoss)
			assert.Equal(t, tc.expectedProfit, remainingProfit, "unexpected remaining profit")
			assert.Equal(t, tc.expectedLoss, remainingLoss, "unexpected remaining loss")
		})
	}
}

func Test_CalculateTax(t *testing.T) {
	tests := []struct {
		name     string
		profit   float64
		expected float64
	}{
		{
			name:     "Simple tax calculation",
			profit:   1000.0,
			expected: 200.0,
		},
		{
			name:     "Zero profit",
			profit:   0.0,
			expected: 0.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := calculateTax(tc.profit)
			assert.Equal(t, tc.expected, result, "unexpected tax")
		})
	}
}

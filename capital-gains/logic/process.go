package logic

import (
	"fmt"
)

// Process calculates the tax for a series of operations based on the rules for gains and losses.
func (l *logic) Process(operations Operations) ([]TaxResult, error) {
	// Resulting tax values for each operation
	taxResults := make([]TaxResult, len(operations))

	// Accumulated loss from previous operations
	var totalLoss float64

	// Weighted average of purchase costs and total quantity of stocks
	var weightedAverage float64
	var totalQuantity int

	// Iterate through each operation and calculate taxes
	for index, op := range operations {
		switch OperationType(op.Operation) {
		case buyOperation:
			// Calculate new weighted average after a buy operation
			weightedAverage, totalQuantity = updateWeightedAverage(weightedAverage, totalQuantity, op.UnitCost, op.Quantity)

			// No tax is paid for buy operations
			taxResults[index] = TaxResult{Tax: 0.00}

		case sellOperation:
			// Calculate gross profit (or loss) for the sell operation
			profit := calculateProfit(op.Total(), float64(op.Quantity)*weightedAverage)

			// Handle accumulated loss for future deductions
			if profit < 0 {
				// Update accumulated loss
				totalLoss = accumulateLoss(totalLoss, profit)
				taxResults[index] = TaxResult{Tax: 0.00}
				totalQuantity -= op.Quantity
				continue
			}

			// Deduct accumulated loss from profit
			profit, totalLoss = deductLoss(profit, totalLoss)

			// Check if the operation is tax-exempt
			if op.TaxFree() {
				taxResults[index] = TaxResult{Tax: 0.00}
				totalQuantity -= op.Quantity
				continue
			}

			// Calculate tax on remaining profit
			tax := calculateTax(profit)
			taxResults[index] = TaxResult{Tax: tax}

			// Update total quantity after the sell
			totalQuantity -= op.Quantity

		default:
			// Invalid operation type
			return nil, fmt.Errorf("invalid operation type %s", op.Operation)
		}
	}

	return taxResults, nil
}

// updateWeightedAverage recalculates the weighted average of purchase costs.
func updateWeightedAverage(weightedAverage float64, totalQuantity int, unitCost float64, quantity int) (float64, int) {
	totalCost := (weightedAverage * float64(totalQuantity)) + (unitCost * float64(quantity))
	totalQuantity += quantity
	return totalCost / float64(totalQuantity), totalQuantity
}

// calculateProfit computes the profit or loss for a sell operation.
func calculateProfit(totalValue, totalCost float64) float64 {
	return totalValue - totalCost
}

// accumulateLoss adds the current loss to the accumulated total loss.
func accumulateLoss(totalLoss, profit float64) float64 {
	return totalLoss + -profit
}

// deductLoss reduces the profit by the accumulated loss, updating the remaining loss.
func deductLoss(profit, totalLoss float64) (float64, float64) {
	if totalLoss > 0 {
		if profit > totalLoss {
			profit -= totalLoss
			totalLoss = 0
		} else {
			totalLoss -= profit
			profit = 0
		}
	}
	return profit, totalLoss
}

// calculateTax computes the tax for the remaining profit.
func calculateTax(profit float64) float64 {
	return round(profit * 0.2)
}

// round rounds a float to two decimal places.
func round(val float64) float64 {
	return float64(int(val*100+0.5)) / 100
}

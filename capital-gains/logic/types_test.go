package logic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OperationTotal(t *testing.T) {
	op := Operation{Operation: "buy", UnitCost: 10.0, Quantity: 100}
	assert.Equal(t, 1000.0, op.Total(), "unexpected total value")
}

func Test_OperationTaxFree(t *testing.T) {
	op := Operation{Operation: "sell", UnitCost: 15.0, Quantity: 1000}
	assert.True(t, op.TaxFree(), "operation should be tax-free")

	op = Operation{Operation: "sell", UnitCost: 30.0, Quantity: 1000}
	assert.False(t, op.TaxFree(), "operation should not be tax-free")
}

func Test_TaxResultMarshalJSON(t *testing.T) {
	tr := TaxResult{Tax: 1234.567}
	expected := `{"tax":1234.57}`
	jsonData, err := tr.MarshalJSON()
	assert.NoError(t, err, "unexpected error during JSON marshaling")
	assert.JSONEq(t, expected, string(jsonData), "unexpected JSON output")
}

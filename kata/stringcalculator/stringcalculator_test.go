package stringcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCalculatorReturnsZeroOnEmptyInput(t *testing.T) {
	res := StringCalculator("")
	assert.Equal(t, 0, res)
}

func TestStringCalculatorReturnsOneOnOneInput(t *testing.T) {
	res := StringCalculator("1")
	assert.Equal(t, 1, res)
}

func TestStringCalculatorWith2Inputs(t *testing.T) {
	res := StringCalculator("1,2")
	assert.Equal(t, 3, res)
}

func TestStringCalcualtorWithMultipleInputs(t *testing.T) {
	res := StringCalculator("1,2,3,4,5")
	assert.Equal(t, 15, res)
}

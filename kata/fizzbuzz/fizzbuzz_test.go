package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzReturnsCorrectString(t *testing.T) {
	inputInt := 10
	outputStr := "10"

	res := FizzBuzz(inputInt)

	assert.Equal(t, outputStr, res)
}

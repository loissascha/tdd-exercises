package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzReturnsCorrectString(t *testing.T) {
	inputInt := 1
	expectedOutput := "1"

	res := FizzBuzz(inputInt)

	assert.Equal(t, expectedOutput, res)
}

func TestFizzBuzzMultipleOf3ReturnsFizz(t *testing.T) {
	inputInt := 6
	expectedOutput := "Fizz"

	res := FizzBuzz(inputInt)

	assert.Equal(t, expectedOutput, res)
}

func TestFizzBuzzMultipleOf5ReturnsBuzz(t *testing.T) {
	inputInt := 20
	expectedOutput := "Buzz"

	res := FizzBuzz(inputInt)

	assert.Equal(t, expectedOutput, res)
}

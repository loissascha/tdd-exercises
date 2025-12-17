package passwordvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordMinLen8(t *testing.T) {
	errors := PasswordValidator("")

	assert.Contains(t, errors, "Password must be at least 8 characters")
}

func TestPasswordMustContain2Numbers(t *testing.T) {
	err := PasswordValidator("TestPassword1")
	noErr := PasswordValidator("TestPassword12#")

	assert.Empty(t, noErr)
	assert.Contains(t, err, "The password must contain at least 2 numbers")
}

func TestPasswordMultiErrors(t *testing.T) {
	errors := PasswordValidator("value")

	assert.Contains(t, errors, "Password must be at least 8 characters")
	assert.Contains(t, errors, "The password must contain at least 2 numbers")
}

func TestPasswordNeedsOneUpper(t *testing.T) {
	errors := PasswordValidator("noupper")

	assert.Contains(t, errors, "The password must contain at least one capital letter")
}

func TestPasswordNeedsSpecialCharacter(t *testing.T) {
	errors := PasswordValidator("nospecial")

	assert.Contains(t, errors, "The password must contain at least one special character")
}

func TestHasUpperValid(t *testing.T) {
	b1 := HasUpper("stRing")
	b2 := HasUpper("String")
	b3 := HasUpper("string")
	b4 := HasUpper("string$")
	b5 := HasUpper("string5")

	assert.True(t, b1)
	assert.True(t, b2)
	assert.False(t, b3)
	assert.False(t, b4)
	assert.False(t, b5)
}

func TestHasSpecialCharacter(t *testing.T) {
	b1 := HasSpecialCharacter("nope")
	b2 := HasSpecialCharacter("Nope")
	b3 := HasSpecialCharacter("Nope22")
	b4 := HasSpecialCharacter("No**")
	b5 := HasSpecialCharacter("No#")

	assert.False(t, b1)
	assert.False(t, b2)
	assert.False(t, b3)
	assert.True(t, b4)
	assert.True(t, b5)
}

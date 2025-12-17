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
	noErr := PasswordValidator("TestPassword12")

	assert.Empty(t, noErr)
	assert.Contains(t, err, "The password must contain at least 2 numbers")
}

func TestPasswordMultiErrors(t *testing.T) {
	errors := PasswordValidator("value")

	assert.Contains(t, errors, "Password must be at least 8 characters")
	assert.Contains(t, errors, "The password must contain at least 2 numbers")
}

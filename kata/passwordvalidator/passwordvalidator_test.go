package passwordvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordMinLen8(t *testing.T) {
	err := PasswordValidator("")

	assert.NotNil(t, err)
	assert.Equal(t, "Password must be at least 8 characters", err.Error())
}

func TestPasswordMustContain2Numbers(t *testing.T) {
	err := PasswordValidator("TestPassword1")
	noErr := PasswordValidator("TestPassword12")

	assert.Nil(t, noErr)
	assert.NotNil(t, err)
	assert.Equal(t, "The password must contain at least 2 numbers", err.Error())
}

package citysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTooShort(t *testing.T) {
	res := Search("a")

	assert.Empty(t, res)
}

package citysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTooShort(t *testing.T) {
	res := Search("a")

	assert.Empty(t, res)
}

func TestSearchVa(t *testing.T) {
	res := Search("Va")

	assert.Contains(t, res, "Valencia")
	assert.Contains(t, res, "Vancouver")
}

func TestSearchva(t *testing.T) {
	res := Search("va")

	assert.Contains(t, res, "Valencia")
	assert.Contains(t, res, "Vancouver")
}

func TestSearchape(t *testing.T) {
	res := Search("ape")

	assert.Contains(t, res, "Budapest")
}

func TestSearchAsterisk(t *testing.T) {
	res := Search("*")

	assert.Equal(t, cities, res)
}

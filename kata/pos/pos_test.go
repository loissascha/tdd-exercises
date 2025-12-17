package pos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	s := NewScanner()
	res1 := s.ScanCode("12345")
	res2 := s.ScanCode("23456")
	res3 := s.ScanCode("99999")
	res4 := s.ScanCode("")

	assert.Equal(t, "$7.25", res1)
	assert.Equal(t, "$12.50", res2)
	assert.Equal(t, "Error: barcode not found", res3)
	assert.Equal(t, "Error: empty barcode", res4)
}

func TestGetTotal(t *testing.T) {
	s := NewScanner()
	s.ScanCode("12345")
	s.ScanCode("23456")
	s.ScanCode("99999")
	s.ScanCode("")

	res := s.GetTotal()

	assert.Equal(t, "$19.75", res)
}

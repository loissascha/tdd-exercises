package pos

import "fmt"

var barcodes = map[string]float64{
	"12345": 7.25,
	"23456": 12.50,
}

type Scanner struct {
	Total float64
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) ScanCode(barcode string) string {
	if barcode == "" {
		return "Error: empty barcode"
	}

	bc, found := barcodes[barcode]
	if !found {
		return "Error: barcode not found"
	}

	s.Total += bc

	return fmt.Sprintf("$%.2f", bc)
}

func (s *Scanner) GetTotal() string {
	return fmt.Sprintf("$%.2f", s.Total)
}

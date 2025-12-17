package passwordvalidator

import (
	"strings"
	"unicode"
)

func PasswordValidator(input string) []string {

	errors := []string{}

	if len(input) < 8 {
		errors = append(errors, "Password must be at least 8 characters")
	}

	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	countNums := 0
	for _, n := range nums {
		countNums += strings.Count(input, n)
	}

	if countNums < 2 {
		errors = append(errors, "The password must contain at least 2 numbers")
	}

	if !HasUpper(input) {
		errors = append(errors, "The password must contain at least one capital letter")
	}

	return errors
}

func HasUpper(input string) bool {
	for _, r := range input {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

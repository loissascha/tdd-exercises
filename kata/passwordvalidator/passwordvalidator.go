package passwordvalidator

import (
	"fmt"
	"strings"
)

func PasswordValidator(input string) error {

	if len(input) < 8 {
		return fmt.Errorf("Password must be at least 8 characters")
	}

	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	countNums := 0
	for _, n := range nums {
		countNums += strings.Count(input, n)
	}

	if countNums < 2 {
		return fmt.Errorf("The password must contain at least 2 numbers")
	}

	return nil
}

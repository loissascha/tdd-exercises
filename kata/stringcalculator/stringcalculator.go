package stringcalculator

import (
	"strconv"
	"strings"
)

func StringCalculator(input string) int {
	split := strings.SplitSeq(input, ",")
	sum := 0
	for s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			sum += 0
		} else {
			sum += n
		}
	}
	return sum
}

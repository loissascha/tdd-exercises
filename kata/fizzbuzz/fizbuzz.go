package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}

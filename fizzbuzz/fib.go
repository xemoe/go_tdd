package fizzbuzz

import (
	"strconv"
)

func Fizzbuzz(s int) string {

	if (s % 3) == 0 {
		return "fizz"
	}

	if (s % 5) == 0 {
		return "buzz"
	}

	return strconv.Itoa(s)
}

package fizzbuzz

import (
	"bytes"
	"strconv"
)

const REASON_FIZZ = 3
const REASON_BUZZ = 5
const WORD_FIZZ = "fizz"
const WORD_BUZZ = "buzz"

func Fizzbuzz() func(int) string {
	return func(s int) string {

		var result bytes.Buffer

		if (s % REASON_FIZZ) == 0 {
			result.WriteString(WORD_FIZZ)
		}

		if (s % REASON_BUZZ) == 0 {
			result.WriteString(WORD_BUZZ)
		}

		if ((s % REASON_FIZZ) != 0) && ((s % REASON_BUZZ) != 0) {
			result.WriteString(strconv.Itoa(s))
		}

		return result.String()
	}
}

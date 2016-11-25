//
// @author Teerapong ladlee
// @update 2016-11-25
//

package fizzbuzz

import (
	"strconv"
)

const REASON_FIZZ = 3
const REASON_BUZZ = 5
const REASON_FIZZBUZZ = 15
const WORD_FIZZ = "fizz"
const WORD_BUZZ = "buzz"
const WORD_FIZZBUZZ = "fizzbuzz"

func Fizzbuzz() func(int) string {
	return func(s int) string {

		if (s % REASON_FIZZBUZZ) == 0 {
			return WORD_FIZZBUZZ
		}

		if (s % REASON_FIZZ) == 0 {
			return WORD_FIZZ
		}

		if (s % REASON_BUZZ) == 0 {
			return WORD_BUZZ
		}

		return strconv.Itoa(s)
	}
}

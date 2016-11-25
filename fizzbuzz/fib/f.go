//
// @author Teerapong ladlee
// @update 2016-11-25
//

package fib

import (
	"strconv"
)

type F []int

const (
	REASON_FIZZ     = 3
	REASON_BUZZ     = 5
	REASON_FIZZBUZZ = 15
	WORD_FIZZ       = "fizz"
	WORD_BUZZ       = "buzz"
	WORD_FIZZBUZZ   = "fizzbuzz"
)

func (f F) Result(s int) string {

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

func (stack *F) Push(x int) {
	*stack = append(*stack, x)
}

func (stack F) All() []string {

	var result []string
	l := len(stack)

	if l == 0 {
		return result
	}

	for _, num := range stack {
		result = append(result, stack.Result(num))
	}

	return result
}

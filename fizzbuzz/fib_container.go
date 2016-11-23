package fizzbuzz

func FizzbuzzContainer() func([]int) []string {
	return func(input []int) []string {

		to_string := Fizzbuzz()
		result := make([]string, len(input))

		for index, num := range input {
			result[index] = to_string(num)
		}

		return result
	}
}

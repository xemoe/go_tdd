package fizzbuzz

func FizzbuzzContainer() func([]int) []string {
	return func(input []int) []string {

		var result []string
		to_string := Fizzbuzz()

		for _, num := range input {
			result = append(result, to_string(num))
		}

		return result
	}
}

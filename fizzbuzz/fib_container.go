package fizzbuzz

func FizzbuzzContainer() func([]int) []string {
	return func(input []int) []string {

		to_string := Fizzbuzz()
		result := make([]string, 0)

		for _, num := range input {
			result = append(result, to_string(num))
		}

		return result
	}
}

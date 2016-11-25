//
// @author Teerapong ladlee
// @update 2016-11-25
//

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

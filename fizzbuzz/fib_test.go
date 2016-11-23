package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzbuzzShouldReturnItValue(t *testing.T) {

	fb_string := Fizzbuzz()
	result := fb_string(1)

	if result != "1" {
		t.Errorf("Fizzbuzz(1) should return 1,%q is returned", result)
	}
}

func TestFizzbuzzShouldReturnFizz(t *testing.T) {

	fb_string := Fizzbuzz()
	result := fb_string(3)

	if result != "fizz" {
		t.Errorf("Fizzbuzz(3) should return fizz, %q is returned", result)
	}
}

func TestFizzbuzzShouldReturnBuzz(t *testing.T) {

	fb_string := Fizzbuzz()
	result := fb_string(5)

	if result != "buzz" {
		t.Errorf("Fizzbuzz(5) should return buzz, %q is returned", result)
	}
}

func TestFizzbuzzShouldReturnFizzbuzz(t *testing.T) {

	fb_string := Fizzbuzz()
	result := fb_string(15)

	if result != "fizzbuzz" {
		t.Errorf("Fizzbuzz(15) should return fizzbuzz, %q is returned", result)
	}
}

func TestFizzBuzzContainerShouldReturnArrayOfExpectedValue(t *testing.T) {

	fb_array := FizzbuzzContainer()
	result := fb_array([]int{1, 2, 3, 4, 5})

	if result == nil {
		t.Errorf("FizzbuzzContainer({...}) should not return nil")
	}

	expected := []string{"1", "2", "fizz", "4", "buzz"}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FizzbuzzContainer({...}) should return expected array, %q is returned", result)
	}
}

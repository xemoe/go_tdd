package fizzbuzz

import (
	"testing"
)

func TestFizzbuzzShouldReturnItValue(t *testing.T) {

	result := Fizzbuzz(1)

	if result != "1" {
		t.Errorf("Fizzbuzz(1) should return 1,%q is returned", result)
	}
}

func TestFizzbuzzShouldReturnFizz(t *testing.T) {

	result := Fizzbuzz(3)

	if result != "fizz" {
		t.Errorf("Fizzbuzz(3) should return fizz, %q is returned", result)
	}
}

func TestFizzbuzzShouldReturnBuzz(t *testing.T) {

	result := Fizzbuzz(5)

	if result != "buzz" {
		t.Errorf("Fizzbuzz(5) should return buzz, %q is returned", result)
	}
}

func TestFizzbuzzShouldReturnFizzbuzz(t *testing.T) {

	result := Fizzbuzz(15)

	if result != "fizzbuzz" {
		t.Errorf("Fizzbuzz(15) should return fizzbuzz, %q is returned", result)
	}
}

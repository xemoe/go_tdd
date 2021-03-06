//
// @author Teerapong ladlee
// @update 2016-11-25
//

package tests

import (
	"github.com/xemoe/go_tdd/fizzbuzz/fib"
	"reflect"
	"testing"
)

func TestResult(t *testing.T) {

	var f fib.F
	var result string

	result = f.Result(1)

	if result != "1" {
		t.Errorf("*F.Result(1) should return 1,%q is returned", result)
	}

	result = f.Result(3)

	if result != "fizz" {
		t.Errorf("*F.Result(3) should return fizz, %q is returned", result)
	}

	result = f.Result(5)

	if result != "buzz" {
		t.Errorf("*F.Result(5) should return buzz, %q is returned", result)
	}

	result = f.Result(15)

	if result != "fizzbuzz" {
		t.Errorf("*F.Result(15) should return fizzbuzz, %q is returned", result)
	}
}

func TestAll(t *testing.T) {

	var f fib.F
	var result []string

	result = f.All()

	if result != nil {
		t.Errorf("*F.All() on empty should return nil")
	}

	f.Push(1)
	f.Push(2)
	f.Push(3)
	f.Push(4)
	f.Push(5)
	f.Push(10)
	f.Push(12)
	f.Push(15)
	f.Push(19)
	f.Push(20)

	result = f.All()

	expected := []string{"1", "2", "fizz", "4", "buzz", "buzz", "fizz", "fizzbuzz", "19", "buzz"}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("*F.All() should return expected array, %q is returned", result)
	}
}

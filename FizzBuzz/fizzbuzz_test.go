package FizzBuzz

import (
	"testing"
)

func TestFizz(t *testing.T) {
	var fizz int = 3
	var buzz int = 5
	sayNumber(fizz, buzz)
}

func TestIsMatch(t *testing.T) {
	var fizz = 3
	result := IsMatch(fizz, 30)
	if !result {
		t.Fail()
	}
}

func TestIsMultiple(t *testing.T) {

	specialNumber := 3
	number := 4
	result := IsMultiple(specialNumber, number)
	if result {
		t.Fail()
	}
	number = 6
	result = IsMultiple(specialNumber, number)
	if !result {
		t.Fail()
	}
}

func TestIsNumberContainSpecialNumber(t *testing.T) {

	specialNumber := 3
	number := 4
	result := IsNumberContainSpecialNumber(specialNumber, number)
	if result {
		t.Fail()
	}
	number = 31
	result = IsNumberContainSpecialNumber(specialNumber, number)
	if !result {
		t.Fail()
	}
}

func TestIsMathTwoSpecialNumber(t *testing.T) {
	var fizz = 3
	var buzz = 5

	result := IsMathTwoSpecialNumber(fizz, buzz, 53)
	if !result {
		t.Fail()
	}

	result = IsMathTwoSpecialNumber(fizz, buzz, 5)
	if result {
		t.Fail()
	}
}

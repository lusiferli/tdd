package FizzBuzz

import (
	"strconv"
	"strings"
)

func sayNumber(fizz int, buzz int) {

}

func IsMatch(fizz int, number int) bool {
	return IsMultiple(fizz, number) || IsNumberContainSpecialNumber(fizz, number)
}

func IsMultiple(specialNumber int, number int) bool {
	return number%specialNumber == 0
}

func IsNumberContainSpecialNumber(specialNumber int, number int) bool {
	return strings.ContainsAny(strconv.Itoa(number), strconv.Itoa(specialNumber))
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsMatch(fizz int, number int) bool {
	return IsMultiple(fizz, number) || IsNumberContainSpecialNumber(fizz, number)
}

func IsMultiple(specialNumber int, number int) bool {
	return number%specialNumber == 0
}

func IsNumberContainSpecialNumber(specialNumber int, number int) bool {
	return strings.ContainsAny(strconv.Itoa(number), strconv.Itoa(specialNumber))
}

func IsMathTwoSpecialNumber(specialNumberOne int, specialNumberTwo int, number int) bool {
	return IsMatch(specialNumberOne, number) && IsMatch(specialNumberTwo, number)
}

func main() {
	fizz := 3
	buzz := 5
	for number := 1; number <= 100; number++ {
		if IsMathTwoSpecialNumber(fizz, buzz, number) {
			fmt.Println("FizzBuzz")
			continue
		}

		if IsMatch(fizz, number) {
			fmt.Println("fizz")
			continue
		}
		if IsMatch(buzz, number) {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(number)
	}
}

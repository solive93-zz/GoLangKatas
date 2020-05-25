package main

import "strconv"

func Calculate(number int) string {
	if isDivisibleByThree(number) && isDivisibleByFive(number) {
		return "FizzBuzz"
	}

	if isDivisibleByThree(number) {
		return "Fizz"
	}

	if isDivisibleByFive(number) {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func isDivisibleByFive(number int) bool {
	return number%5 == 0
}

func isDivisibleByThree(number int) bool {
	return number%3 == 0
}

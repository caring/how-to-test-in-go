package fizzbuzz

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) (string, error) {
	if n < 0 {
		panic(fmt.Sprint("n is negative. got = ", n))
	}
	if n == 0 {
		return "", fmt.Errorf("too much fizzbuzzery")
	}

	switch {
	case n%15 == 0:
		return "fizzbuzz", nil

	case n%5 == 0:
		return "buzz", nil

	case n%3 == 0:
		return "fizz", nil

	default:
		return strconv.Itoa(int(n)), nil
	}
}

package main

import (
	"math"
	"strconv"
)

func main() {}

func solve(base, maxLength int) []string {
	xs := make([]string, 0, maxLength)

	var maxDigits int
	for digit := 1; ; digit++ {
		if uint(digit)*pow(base-1, digit) < pow(base, digit-1) {
			maxDigits = digit
			break
		}
	}
	max := pow(base, maxDigits)

	for i := base; uint(i) <= max; i++ {
		x := strconv.FormatInt(int64(i), base)
		if !narcissist(base, x) {
			continue
		}

		xs = append(xs, x)
		if maxLength <= len(xs) {
			break
		}
	}

	return xs
}

func narcissist(base int, x string) bool {
	var sum uint

	for _, rawDigit := range x {
		digit, err := strconv.ParseInt(string(rawDigit), base, 64)
		if err != nil {
			return false
		}
		sum += pow(int(digit), len(x))
	}

	return strconv.FormatInt(int64(sum), base) == x
}

func pow(x, y int) uint {
	return uint(math.Pow(float64(x), float64(y)))
}

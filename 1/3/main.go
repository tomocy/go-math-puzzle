package main

import (
	"strings"
)

func main() {}

func solve(n int) int {
	var patterns int
	for i := 1; i <= 3999; i++ {
		roman := m(i)
		if len(roman) != n {
			continue
		}
		patterns++
	}

	return patterns
}

func m(n int) string {
	return roman(n, "M", 1000, cm)
}

func cm(n int) string {
	return roman(n, "CM", 900, d)
}

func d(n int) string {
	return roman(n, "D", 500, cd)
}

func cd(n int) string {
	return roman(n, "CD", 400, c)
}

func c(n int) string {
	return roman(n, "C", 100, xc)
}

func xc(n int) string {
	return roman(n, "XC", 90, l)
}

func l(n int) string {
	return roman(n, "L", 50, xl)
}

func xl(n int) string {
	return roman(n, "XL", 40, x)
}

func x(n int) string {
	return roman(n, "X", 10, ix)
}

func ix(n int) string {
	return roman(n, "IX", 9, v)
}

func v(n int) string {
	return roman(n, "V", 5, iv)
}

func iv(n int) string {
	return roman(n, "IV", 4, i)
}

func i(n int) string {
	return roman(n, "I", 1, func(int) string { return "" })
}

func roman(n int, symbol string, base int, next func(int) string) string {
	return strings.Repeat(symbol, n/base) + next(n%base)
}

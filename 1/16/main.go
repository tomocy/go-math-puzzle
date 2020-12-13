package main

func main() {}

func solve(n int) int {
	var x int

	for i := 1; float64(i) < float64(n)/2; i++ {
		if gcd(i, n-i) == 1 {
			x++
		}
	}

	return x
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

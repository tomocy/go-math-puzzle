package main

func main() {}

func solve(n int) int {
	cache := make(map[int]int)
	var x int

	for i := 2; i < n; i++ {
		if n%i != 0 {
			continue
		}

		x += seats(cache, i)
	}

	return x
}

func seats(cache map[int]int, w int) int {
	if x, ok := cache[w]; ok {
		return x
	}

	var x int

	if w < 6 {
		x++
	}

	for i := 2; i < 6; i++ {
		rest := w - i*2
		if rest < 0 {
			break
		}

		if rest == 0 {
			x++
			continue
		}

		if rest > 1 {
			x += seats(cache, w-i*2)
		}
	}

	cache[w] = x
	return x
}

package main

func main() {}

func solve(length, n int) int {
	nReverseds := make([][][]int, n+1)

	nReverseds[0] = generateStoppedPatterns(length)

	for i, reverseds := range nReverseds {
		if i == len(nReverseds)-1 {
			break
		}

		for _, reversed := range reverseds {
			for j, v := range reversed {
				if v == 1 || v != j+1 {
					continue
				}

				nReverseds[i+1] = append(nReverseds[i+1], reverse(reversed, j))
			}
		}
	}

	return len(nReverseds[n])
}

func generateStoppedPatterns(n int) [][]int {
	vs := make([]int, n-1)
	for i := range vs {
		vs[i] = i + 2
	}
	permutated := permutate(vs)

	var stopped [][]int
	for _, p := range permutated {
		stopped = append(stopped, append([]int{1}, p...))
	}

	return stopped
}

func permutate(vs []int) [][]int {
	xs := make([][]int, 0, factorial(len(vs)))
	permutateK(&xs, vs, len(vs))
	return xs
}

func permutateK(dst *[][]int, vs []int, k int) {
	if k == 1 {
		*dst = append(*dst, clone(vs))
		return
	}

	for i := 0; i < k; i++ {
		permutateK(dst, vs, k-1)
		j := k % 2 * i
		vs[j], vs[k-1] = vs[k-1], vs[j]
	}
}

func clone(nums []int) []int {
	return append([]int{}, nums...)
}

func factorial(n int) int {
	x := 1
	for i := n; i > 0; i-- {
		x *= i
	}
	return x
}

func reverse(vs []int, i int) []int {
	reversed := make([]int, len(vs))
	for asc, des := 0, i; des >= 0; asc, des = asc+1, des-1 {
		reversed[asc] = vs[des]
	}
	copy(reversed[i+1:], vs[i+1:])

	return reversed
}

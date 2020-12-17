package main

func main() {}

func solve(n, start, end int) int {
	if start == end {
		return 1
	}
	if start < end {
		return count(n-end, end-2)
	}
	return count(end-1, n-end-1)
}

func count(forward, backward int) int {
	if forward == 0 {
		return 1
	}

	return 1 + forward*count(backward, forward-1)
}

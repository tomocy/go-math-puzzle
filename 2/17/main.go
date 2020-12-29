package main

func main() {}

func solve(given, stock int) int {
	interX, _ := intersection(given, stock)
	patterns := smallChoco(0, given, stock) + 1
	if interX > 0 {
		patterns -= interX
	}
	return patterns
}

func intersection(given, stock int) (int, int) {
	return 2*given - stock, -given + stock
}

func smallChoco(medium, given, stock int) int {
	return (3*given - stock - medium) / 2
}

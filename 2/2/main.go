package main

func main() {}

func solve(n int) int {
	var x int

	for i := uint(1); i < 1<<n; i++ {
		x += escape(i)
	}

	return x
}

func escape(n uint) int {
	if n == 0 {
		return 0
	}

	n = move(n)
	return 1 + escape(n)
}

func move(n uint) uint {
	empty := ^n
	movable := (empty << 1) + 1
	unmoved, moved := (n & ^movable), (n>>1)&empty
	return unmoved | moved
}

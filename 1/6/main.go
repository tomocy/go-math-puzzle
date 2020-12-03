package main

func main() {}

func solve(n, maxLength int) int {
	var x int

	for shorter := 1; shorter <= maxLength; shorter++ {
		for longer := shorter; longer <= maxLength; longer++ {
			squares := cutToSquare(shorter, longer)
			if squares != n {
				continue
			}
			x++
		}
	}

	return x
}

func cutToSquare(width, height int) int {
	if width == height {
		return 1
	}

	shorter, longer := width, height
	if longer < shorter {
		shorter, longer = longer, shorter
	}

	squares, rest := longer/shorter, longer%shorter
	if rest != 0 {
		squares += cutToSquare(rest, shorter)
	}

	return squares
}

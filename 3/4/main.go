package main

func main() {}

func solve(w, h int) int {
	outside, inside := (w+h)*2, (w-1)*(h-1)*2
	if w%2 == 1 || h%2 == 1 {
		return outside + inside
	}
	return outside + inside - 2
}

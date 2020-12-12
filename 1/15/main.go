package main

func main() {}

func solve(n, k int) int {
	prev, now := 0, 1
	for k > 1 {
		if now == prev*2 || now == prev*2+1 {
			if now*2 <= n {
				prev, now, k = now, now*2, k-1
				continue
			}

			prev, now = now, now/2
			continue
		}

		if now*2 == prev {
			if now*2+1 <= n {
				prev, now, k = now, now*2+1, k-1
				continue
			}

			prev, now = now, now/2
			continue
		}

		if now*2+1 == prev {
			prev, now = now, now/2
			continue
		}
	}

	return now
}

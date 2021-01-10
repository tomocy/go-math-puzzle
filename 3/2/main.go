package main

func main() {}

func solve(n int) int {
	return sail(make([]int, n+1), 0, 0)
}

func sail(levels []int, pos, level int) int {
	levels[pos] = level

	if pos == len(levels)-1 {
		return distance(levels)
	}

	var max int

	if level > 0 {
		if guessed := sail(append([]int{}, levels...), pos+1, level-1); max < guessed {
			max = guessed
		}
	}
	if nextPos, nextLevel := pos+1, level+1; nextPos+nextLevel <= len(levels)-1 {
		if guessed := sail(append([]int{}, levels...), nextPos, nextLevel); max < guessed {
			max = guessed
		}
	}

	return max
}

func distance(levels []int) int {
	queue := []sailCtx{
		{
			left: 0, right: len(levels) - 1,
		},
	}
	dist := make(map[sailCtx]int)

	for len(queue) > 0 {
		ctx := queue[0]
		queue = queue[1:]

		if ctx.left == ctx.right {
			return dist[ctx]
		}

		for _, next := range ctx.next() {
			if next.left < 0 || next.right > len(levels)-1 || next.left > next.right {
				continue
			}
			if levels[next.left] != levels[next.right] {
				continue
			}
			if _, ok := dist[next]; ok {
				continue
			}

			queue = append(queue, next)
			dist[next] = dist[ctx] + 2
		}
	}

	return 0
}

type sailCtx struct {
	left, right int
}

func (c sailCtx) next() []sailCtx {
	return []sailCtx{
		{
			left:  c.left - 1,
			right: c.right - 1,
		},
		{
			left:  c.left + 1,
			right: c.right - 1,
		},
		{
			left:  c.left - 1,
			right: c.right + 1,
		},
		{
			left:  c.left + 1,
			right: c.right + 1,
		},
	}
}

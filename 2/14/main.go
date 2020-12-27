package main

func main() {}

func solve(cards, limit int) int {
	return take(make(map[takeCtx]int), cards, cards, limit, 0, turnSecond)
}

func take(cache map[takeCtx]int, cards, remain, limit, firstCards int, prevTurn turn) int {
	ctx := takeCtx{
		cards:      cards,
		remain:     remain,
		limit:      limit,
		firstCards: firstCards,
		prevTurn:   prevTurn,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if remain == 0 {
		if prevTurn == turnFirst && firstCards > cards-firstCards {
			cache[ctx] = 1
		} else {
			cache[ctx] = 0
		}
	} else {
		for i := 1; i <= limit; i++ {
			if i > remain {
				continue
			}
			if prevTurn == turnFirst {
				cache[ctx] += take(cache, cards, remain-i, limit, firstCards, turnSecond)
			} else {
				cache[ctx] += take(cache, cards, remain-i, limit, firstCards+i, turnFirst)
			}
		}
	}

	return cache[ctx]
}

type takeCtx struct {
	cards      int
	remain     int
	limit      int
	firstCards int
	prevTurn   turn
}

const (
	turnFirst turn = iota
	turnSecond
)

type turn int

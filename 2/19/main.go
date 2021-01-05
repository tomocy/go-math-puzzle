package main

const (
	amount                = 3000
	b1000sForAmount       = amount / 1000
	changedB1000sForB5000 = (5000 - amount) / 1000
)

func main() {}

func solve(members int) int {
	return organize(make(map[organizeCtx]int), 0, members)
}

func organize(cache map[organizeCtx]int, bills, members int) int {
	ctx := organizeCtx{
		bills:   bills,
		members: members,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if members == 0 {
		cache[ctx] = 1
		return cache[ctx]
	}

	cache[ctx] += organize(cache, bills+b1000sForAmount, members-1)
	if changed := bills - changedB1000sForB5000; changed >= 0 {
		cache[ctx] += organize(cache, changed, members-1)
	}

	return cache[ctx]
}

type organizeCtx struct {
	bills   int
	members int
}

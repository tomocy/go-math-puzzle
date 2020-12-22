package main

func main() {}

func solve(row int) int {
	cache := new(cache)
	return solveWithCache(cache, row)
}

func solveWithCache(cache *cache, row int) int {
	if x, ok := cache.get(row); ok {
		return x
	}

	x := twoSeats(row) * threeSeats(row)

	for i := 0; i < row; i++ {
		x += twoSeats(i) * threeSeats(i) * solveWithCache(cache, row-i-1)
	}

	cache.set(row, x)
	return x
}

func twoSeats(row int) int {
	if row < 0 {
		return 0
	}
	if row == 0 {
		return 1
	}

	return 2*twoSeats(row-1) + twoSeats(row-2)
}

func threeSeats(row int) int {
	if row < 0 {
		return 0
	}
	if row == 0 {
		return 1
	}

	return 4*threeSeats(row-1) + threeSeats(row-2)
}

type cache struct {
	data     map[int]int
	counts   map[int]int
	hits     map[int]int
	doIgnore bool
}

func (c *cache) get(row int) (int, bool) {
	c.init()

	c.counts[row]++
	v, ok := c.data[row]
	if c.doUse() && ok {
		c.hits[row]++
	}
	return v, c.doUse() && ok
}

func (c cache) doUse() bool {
	return !c.doIgnore
}

func (c *cache) set(row, v int) {
	c.init()

	c.data[row] = v
}

func (c *cache) hitRate(row int) float64 {
	c.init()

	return float64(c.hits[row]) / float64(c.counts[row])
}

func (c *cache) init() {
	if c.data == nil {
		c.data = make(map[int]int)
	}
	if c.counts == nil {
		c.counts = make(map[int]int)
	}
	if c.hits == nil {
		c.hits = make(map[int]int)
	}
}

package main

import (
	"math"
)

func main() {}

func solve(n int) int {
	pow := int(math.Pow10(n))
	pi := int(math.Pi * float64(pow))

	var minDen int
	for den := 1; ; den++ {
		if pi*den/pow == (pi+1)*den/pow {
			continue
		}

		if (pi+1)*den%pow == 0 {
			continue
		}

		minDen = den
		break
	}

	return minDen
}

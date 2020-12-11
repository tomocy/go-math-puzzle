package main

var billsCoins = [...]int{10000, 5000, 2000, 1000, 500, 100, 50, 10, 5, 1}

func main() {}

func solve(n int) int {
	remain := n

	used := make([]int, len(billsCoins))
	for i := len(used) - 1; i >= 0; i-- {
		if remain-billsCoins[i] < 0 {
			break
		}

		used[i] += 1
		remain -= billsCoins[i]
	}

	for i, billCoin := range billsCoins {
		used[i] += remain / billCoin
		remain %= billCoin
	}

	for i, billCoin := range billsCoins {
		if i == len(billsCoins)-2 {
			break
		}

		if used[i] == 0 && billsCoins[i+1]*used[i+1] >= billCoin {
			used[i], used[i+1] = 1, billCoin/billsCoins[i+1]
		}
	}

	return sum(used)
}

func sum(vs []int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}
	return sum
}

package main

func main() {}

func solve(n int) int {
	var x int

	row := rowOfPascalTriangle(n)
	for _, amount := range row {
		x += tenThousandsYen(amount)
	}

	return x
}

func rowOfPascalTriangle(n int) []int {
	rows := make([][]int, n+1)
	for i := range rows {
		if i == 0 {
			rows[i] = []int{1}
			continue
		}
		if i == 1 {
			rows[i] = []int{1, 1}
			continue
		}

		prevRow := rows[i-1]

		rows[i] = append(rows[i], 1)
		for left, right := 0, 1; right < len(prevRow); left, right = right, right+1 {
			rows[i] = append(rows[i], prevRow[left]+prevRow[right])
		}
		rows[i] = append(rows[i], 1)
	}

	return rows[n]
}

func tenThousandsYen(amount int) int {
	return billsCoins(amount, 10000, fiveThousandsYen)
}

func fiveThousandsYen(amount int) int {
	return billsCoins(amount, 5000, twoThousandsYen)
}

func twoThousandsYen(amount int) int {
	return billsCoins(amount, 2000, oneThousandYen)
}

func oneThousandYen(amount int) int {
	return billsCoins(amount, 1000, fiveHundredsYen)
}

func fiveHundredsYen(amount int) int {
	return billsCoins(amount, 500, oneHundredYen)
}

func oneHundredYen(amount int) int {
	return billsCoins(amount, 100, fiftyYen)
}

func fiftyYen(amount int) int {
	return billsCoins(amount, 50, tenYen)
}

func tenYen(amount int) int {
	return billsCoins(amount, 10, fiveYen)
}

func fiveYen(amount int) int {
	return billsCoins(amount, 5, yen)
}

func yen(amount int) int {
	return amount
}

func billsCoins(amount, v int, next func(int) int) int {
	return amount/v + next(amount%v)
}

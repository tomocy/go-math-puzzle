package main

func main() {}

func solve(n int) int {
	var x int

	for hour := 0; hour < 24; hour++ {
		for minute := 0; minute < 60; minute++ {
			for second := 0; second < 60; second++ {
				amount := amountOfTime(hour, minute, second)
				if amount != n {
					continue
				}
				x++
			}
		}
	}

	return x
}

func amountOfTime(hour, minute, second int) int {
	return amountOfTimeUnit(hour) + amountOfTimeUnit(minute) + amountOfTimeUnit(second)
}

func amountOfTimeUnit(unit int) int {
	return amounts[unit/10] + amounts[unit%10]
}

var amounts = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

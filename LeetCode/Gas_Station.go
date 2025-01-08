package main

// Неэффективное решение
func canCompleteCircuit(gas []int, cost []int) int {
	payment := make([]int, len(gas))
	ans := -1
	for i := 0; i < len(gas); i++ {
		payment[i] = gas[i] - cost[i]
	}
	left := 0
	for i := 0; i < len(gas); i++ {
		left += gas[i] - cost[i]
	}
	if left < 0 {
		return ans
	}
	for i := 0; i < len(gas); i++ {
		if payment[i] >= 0 {
			left = 0
			for j := i; j < i+len(payment); j++ {
				left += payment[j%len(payment)]
				if left < 0 {
					break
				}
			}
			if left >= 0 {
				ans = i
				break
			}
		}
	}
	return ans
}

// Эффективное решение
func canCompleteCircuit_v2(gas []int, cost []int) int {
	fuelleft, globalfuelleft, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		globalfuelleft += gas[i] - cost[i]
		fuelleft += gas[i] - cost[i]
		if fuelleft < 0 {
			start = i + 1
			fuelleft = 0
		}
	}
	if globalfuelleft < 0 {
		return -1
	}
	return start
}

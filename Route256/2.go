package main

import (
	"strconv"
)

func tables(n int) int {
	nStr := strconv.Itoa(n)
	count := 0
	if len(nStr) == 1 {
		count = n
		count++
		return count
	}
	count = 9
	count += 10 * (len(nStr) - 1)
	curNum := 1
	ncopy := n
	for ncopy != 0 {
		curNum *= 10
		ncopy = ncopy / 10
	}
	curNum--
	minusS := make([]byte, len(nStr))
	for i := 0; i < len(nStr); i++ {
		minusS[i] = '1'
	}
	minus, _ := strconv.Atoi(string(minusS))
	for i := 0; curNum > n; i++ {
		count--
		curNum -= minus
	}

	return count
}

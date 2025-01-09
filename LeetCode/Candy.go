package main

import "fmt"

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	total_candys := 0
	candys[0] = 1
	total_candys++
	for i := 1; i < len(candys); i++ {
		candys[i] = 1
		total_candys++
		if ratings[i] > ratings[i-1] {
			buf := candys[i]
			candys[i] = candys[i-1] + 1
			total_candys += candys[i] - buf
		}
	}
	fmt.Println(candys)
	for i := len(candys) - 2; i > -1; i-- {
		if candys[i] <= candys[i+1] && ratings[i] > ratings[i+1] {
			buf := candys[i]
			candys[i] = candys[i+1] + 1
			total_candys += candys[i] - buf
		}
	}
	fmt.Println(candys)
	return total_candys
}

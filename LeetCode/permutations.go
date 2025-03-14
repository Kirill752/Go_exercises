package main

import (
	"slices"
)

// S1: (1)
// S2: (2, 1) (1, 2)
// S3: (3,2,1) (2,1,3) (2, 1, 3) (3,1,2), (1, 3, 2), (1, 2, 3)
// ....

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func add(num int, nums []int) [][]int {
	res := [][]int{}
	for i := 0; i <= len(nums); i++ {
		c := make([]int, 0, len(nums))
		c = append(c, nums...)
		res = append(res, slices.Insert(c, i, num))
	}
	return res
}

func permute(input []int) [][]int {
	res := make([][]int, 0, fact(len(input)))
	res = append(res, []int{input[0]})
	for _, num := range input[1:] {
		for _, nums := range res {
			res = append(res, add(num, nums)...)
			res = res[1:]
		}
	}
	return res
}

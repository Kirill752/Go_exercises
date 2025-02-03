package main

func twoSumSorted(nums []int, target int) []int {
	L, R := 0, len(nums)-1
	res := make([]int, 2)
	for L < R {
		switch {
		case nums[L]+nums[R] > target:
			R--
		case nums[L]+nums[R] < target:
			L++
		default:
			res[0] = L + 1
			res[1] = R + 1
			return res
		}
	}
	return res
}

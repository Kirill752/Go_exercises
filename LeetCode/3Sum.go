package main

import "slices"

func twoSumSrt(nums []int, target int) (result [][]int) {
	L, R := 0, len(nums)-1
	resMap := make(map[[2]int]bool)
	for L < R {
		switch {
		case nums[L]+nums[R] > target:
			R--
		case nums[L]+nums[R] < target:
			L++
		default:
			res := [2]int{nums[L], nums[R]}
			if _, ok := resMap[res]; !ok {
				result = append(result, []int{nums[L], nums[R]})
			}
			resMap[res] = true
			L++
		}
	}
	return
}

func threeSum(nums []int) (result [][]int) {
	slices.Sort(nums)
	for L := 0; L < len(nums)-1; L++ {
		if L > 0 && nums[L] == nums[L-1] {
			continue
		}
		target := 0 - nums[L]
		good := twoSumSrt(nums[L+1:], target)
		for i := 0; i < len(good); i++ {
			result = append(result, []int{nums[L], good[i][0], good[i][1]})
		}
	}
	return
}

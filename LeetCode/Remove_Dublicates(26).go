package main

import "fmt"

func removeDuplicates(nums []int) int {
	R := 0
	uniq := []int{nums[0]}
	for L := 0; L < len(nums); L++ {
		if uniq[len(uniq)-1] != nums[L] {
			uniq = append(uniq, nums[L])
		}
		for R < len(nums) && nums[R] == nums[L] {
			R++
		}
	}
	for i := 0; i < len(uniq); i++ {
		nums[i] = uniq[i]
	}
	fmt.Println(nums)
	return len(uniq)
}
func removeDuplicates_2(nums []int) int {
	res := 1
	uniq := nums[0]
	for L := 0; L < len(nums); L++ {
		if uniq != nums[L] {
			uniq = nums[L]
			nums[res] = uniq
			res++
		}
	}
	return res
}

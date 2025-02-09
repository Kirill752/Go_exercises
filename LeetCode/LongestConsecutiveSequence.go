package main

import "sort"

func longestConsecutive(nums []int) int {
	sequence := make(map[int]int)
	maxLen := 0
	if len(nums) > 0 {
		maxLen = 1
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i++ {
		if _, ok := sequence[nums[i]-1]; !ok {
			sequence[nums[i]] = 1
		} else {
			sequence[nums[i]] = sequence[nums[i]-1] + 1
			maxLen = max(maxLen, sequence[nums[i]])
		}
	}
	return maxLen
}

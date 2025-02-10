package main

import "fmt"

func summaryRanges(nums []int) (ranges []string) {
	if len(nums) == 0 {
		return
	}
	if len(nums) == 1 {
		ranges = append(ranges, fmt.Sprintf("%d", nums[0]))
		return
	}
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		for ; j < len(nums) && nums[j] == nums[j-1]+1; j++ {
		}
		j--
		if nums[j] != nums[i] {
			ranges = append(ranges, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		} else {
			ranges = append(ranges, fmt.Sprintf("%d", nums[i]))
		}
		i = j
	}
	return
}

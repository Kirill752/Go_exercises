package main

func containsNearbyDuplicate(nums []int, k int) bool {
	valIndx := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := valIndx[nums[i]]; !ok {
			valIndx[nums[i]] = i
		} else {
			if i-valIndx[nums[i]]+1 <= k {
				return true
			} else {
				valIndx[nums[i]] = i
			}
		}
	}
	return false
}

package main

func minSubArrayLen(target int, nums []int) int {
	L := 0
	if len(nums) == 1 {
		if nums[L] >= target {
			return 1
		}
		return 0
	}
	R := 1
	currentSum := nums[L]
	minLen := 1000001
	for ; R < len(nums); R++ {
		if currentSum < target {
			currentSum += nums[R]
		} else {
			R--
			for currentSum >= target {
				minLen = min(minLen, (R-L)+1)
				currentSum -= nums[L]
				L++
			}
		}
	}
	R--
	for currentSum >= target {
		minLen = min(minLen, (R-L)+1)
		currentSum -= nums[L]
		L++
	}
	if minLen > len(nums) {
		return 0
	}
	return minLen
}

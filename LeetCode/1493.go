package main

func longestSubarray(nums []int) int {
	prefix := make([]int, len(nums)+1)
	result := 0
	prefix[0] = 0
	nonZERO := true
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
		if nums[i] == 0 {
			nonZERO = false
		}
	}
	if nonZERO == true {
		return len(nums) - 1
	}
	L, R := 0, 0
	current_sum := 0
	for R < len(nums) {
		for (R - L) > (prefix[R+1] - prefix[L]) {
			L++
			current_sum = prefix[R+1] - prefix[L]
			result = max(result, current_sum)
			if L > R {
				break
			}
		}
		current_sum = prefix[R+1] - prefix[L]
		result = max(result, current_sum)
		R++
	}
	return result
}

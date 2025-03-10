package main

//	L
//
// [1,1,0,0,1,1,1,0,1]
//
//	R
func longestSubarray(nums []int) int {
	zeroCounter := 0
	resLen := 0
	L, R := 0, 0
	// Пропуск нулей в начале
	for ; L < len(nums) && nums[L] == 0; L++ {
	}
	R = L
	for R < len(nums) {
		if nums[R] != 1 {
			zeroCounter++
			if zeroCounter > 1 {
				for ; zeroCounter > 1; L++ {
					if nums[L] == 0 {
						zeroCounter--
					}
				}
			}
		}
		resLen = max(resLen, R-L+1)
		R++
	}
	if resLen == 0 {
		return resLen
	}
	return resLen - 1
}

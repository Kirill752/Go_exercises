package main

func productExceptSelf(nums []int) []int {
	pref := make([]int, len(nums))
	postf := make([]int, len(nums))
	answer := make([]int, len(nums))
	pref[0] = nums[0]
	postf[len(postf)-1] = nums[len(nums)-1]
	for i := 1; i < len(pref); i++ {
		pref[i] = pref[i-1] * nums[i]
	}
	for i := len(postf) - 2; i > -1; i-- {
		postf[i] = postf[i+1] * nums[i]
	}
	answer[0] = postf[1]
	for i := 1; i < len(answer)-1; i++ {
		answer[i] = pref[i-1] * postf[i+1]
	}
	answer[len(answer)-1] = pref[len(answer)-2]
	return answer
}

package main

func rotate(nums []int, k int) {
	nums_copy := nums
	for i := 0; i < len(nums); i++ {
		nums_copy[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums_copy[i]
	}
}
func rotate_InPlace(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

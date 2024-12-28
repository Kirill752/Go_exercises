package main

func removeElement(nums []int, val int) int {
	var k int = 0
	for i, v := range nums {
		if v == val {
			nums[i] = -1 // так как 0 <= nums[i] <= 50
		} else {
			k++
		}
	}

	return k
}

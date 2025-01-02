package main

func removeDuplicates_II(nums []int) int {
	k := len(nums)
	uniq := nums[0]
	uniq_count := 1
	// fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == uniq {
			uniq_count++
			if uniq_count > 2 {
				nums[i] = -10001
				k--
			}
		} else {
			uniq = nums[i]
			uniq_count = 1
		}
	}
	// fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == -10001 {
			j := i
			for ; j < len(nums) && (nums[j] == -10001); j++ {
			}
			swap(nums, i, j)
		}
	}
	// fmt.Println(nums)
	return k
}

package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0], res[1] = i, j
			}
		}
	}
	return res
}
func twoSum_2(nums []int, target int) []int {
	res_map := make(map[int]int)
	res := make([]int, 2)
	for i, v := range nums {
		if value, exist := res_map[(target - v)]; exist {
			res[0] = value
			res[1] = i
		} else {
			res_map[v] = i
		}
	}
	return res
}

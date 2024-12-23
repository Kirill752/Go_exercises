package main

import "fmt"

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
func main() {
	input := []int{2, 7, 11, 15}
	result := twoSum(input, 9)
	fmt.Println(result)
}

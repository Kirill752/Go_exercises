package main

func subarraySum(nums []int, k int) int {
	res := 0
	prefix := 0
	mapPrefix := make(map[int]int)
	mapPrefix[0]++
	for _, num := range nums {
		prefix += num
		if count, ok := mapPrefix[prefix-k]; ok {
			res += count
		}
		mapPrefix[prefix]++
	}
	return res
}

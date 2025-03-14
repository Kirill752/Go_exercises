package main

func combByK(n, k, idx int, curr []int, result *[][]int) {
	if idx == k-1 {
		c := make([]int, 0, len(curr))
		c = append(c, curr...)
		*result = append(*result, c)
		return
	}
	for i := curr[idx]; i < n; i++ {
		combByK(n, k, idx+1, append(curr, i+1), result)
	}
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		curr := []int{i + 1}
		combByK(n, k, 0, curr, &res)
	}
	return res
}

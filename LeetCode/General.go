package main

func reverse(nums []int, i int, j int) {
	L, R := min(i, j), max(i, j)
	for L < R {
		swap(nums, L, R)
		L++
		R--
	}
}
func swap(massive []int, i, j int) {
	if (i < len(massive)) && (j < len(massive)) {
		buf := massive[i]
		massive[i] = massive[j]
		massive[j] = buf
	}
}

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

// Реализация быстрой сортировки
func partion(nums []int, l int, r int) int {
	x := nums[r]
	less := l
	for i := l; i < r; i++ {
		if nums[i] <= x {
			swap(nums, i, less)
			less++
		}
	}
	swap(nums, less, r)
	return less
}

func qsortImpl(nums []int, begin int, end int) {
	if begin < end {
		ref := partion(nums, begin, end)
		qsortImpl(nums, begin, ref-1)
		qsortImpl(nums, ref+1, end)
	}
}
func qsort(nums []int) {
	if len(nums) > 0 {
		qsortImpl(nums, 0, len(nums)-1)
	}
}

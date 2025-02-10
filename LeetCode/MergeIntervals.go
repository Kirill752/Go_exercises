package main

import "slices"

func merge(intervals [][]int) (mergedIntervals [][]int) {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		if a[0] == b[0] && a[1] == b[1] {
			return 0
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		j := i
		for ; j < len(intervals); j++ {
			if intervals[j][0] < start && intervals[j][1] < start {
				break
			}
			if intervals[j][0] > end && intervals[j][1] > end {
				break
			}
			start = min(start, intervals[j][0])
			end = max(end, intervals[j][1])
		}
		j--
		mergedIntervals = append(mergedIntervals, []int{start, end})
		if j+1 < len(intervals) {
			start, end = intervals[j+1][0], intervals[j+1][1]
		}
		i = j
	}
	return
}

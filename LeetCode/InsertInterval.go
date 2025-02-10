package main

import "slices"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	start, end := newInterval[0], newInterval[1]
	i := 0
	for ; i < len(intervals); i++ {
		if start < intervals[i][0] && end < intervals[i][0] {
			continue
		}
		if start > intervals[i][1] && end > intervals[i][1] {
			continue
		}
		break
	}
	if i == len(intervals) {
		i--
	}
	insertInd := i
	if start >= intervals[i][0] {
		insertInd++
	}
	if insertInd >= len(intervals)-1 {
		intervals = append(intervals, newInterval)
	} else {
		intervals = slices.Insert(intervals, insertInd, newInterval)
	}
	return merge(intervals)
}

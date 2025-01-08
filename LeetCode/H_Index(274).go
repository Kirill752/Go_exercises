package main

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool { return citations[i] < citations[j] })
	hirsh := 0
	for i, v := range citations {
		if len(citations)-i >= v {
			hirsh = max(hirsh, v)
		} else {
			hirsh = max(hirsh, len(citations)-i)
		}
	}
	return hirsh
}

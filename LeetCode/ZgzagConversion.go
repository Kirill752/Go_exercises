package main

import "sort"

func convertZigZag(s string, numRows int) string {
	wordZigZag := make(map[int][]byte)
	up := true
	j := 0
	for i := 0; i < len(s); i++ {
		if up {
			wordZigZag[j] = append(wordZigZag[j], s[i])
			j++
			if j == numRows {
				up = false
				j--
			}
		} else {
			j--
			if j < 0 {
				j = 0
			}
			wordZigZag[j] = append(wordZigZag[j], s[i])
			if j == 0 {
				up = true
				j++
			}
		}
	}
	res := []byte{}
	keys := make([]int, len(wordZigZag))
	for k := range wordZigZag {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return i < j })
	for k := range keys {
		res = append(res, wordZigZag[k]...)
	}
	return string(res)
}

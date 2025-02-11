package main

import "sort"

func isPalindrom(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for i <= j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

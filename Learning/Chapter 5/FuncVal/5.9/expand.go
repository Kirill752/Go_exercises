package main

import (
	"slices"
)

func expand(s string, f func(string) string) string {
	// sptr := unsafe.StringData(s)
	sSlice := []rune(s)
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] == '$' {
			start := i
			end := i
			for ; end < len(sSlice) && sSlice[end] != ' ' && sSlice[end] != '\n'; end++ {
			}
			newText := []rune(f(string(sSlice[start:end])))
			sSlice = slices.Delete(sSlice, start, end)
			sSlice = slices.Insert(sSlice, start, newText...)
		}
	}
	return string(sSlice)
}

package main

import "slices"

func reverseWords(s string) string {
	chrs := []byte(s)
	i := len(chrs) - 1
	sRev := []byte{}
	L := 0
	for ; chrs[L] == ' '; L++ {
	}
	for ; i > L-1; i-- {
		word := []byte{' '}
		for ; i > L-1 && chrs[i] == ' '; i-- {
		}
		for ; i > L-1 && chrs[i] != ' '; i-- {
			word = append(word, chrs[i])
		}
		slices.Reverse(word)
		sRev = append(sRev, word...)
	}
	return string(sRev[:len(sRev)-1])
}

package main

import (
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	var compress strings.Builder
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		l := 0
		j := i
		for ; j < len(chars) && chars[j] == char; j, l = j+1, l+1 {
		}
		compress.WriteByte(char)
		if l > 1 {
			compress.WriteString(strconv.Itoa(l))
		}
		i = j - 1
	}
	chars = chars[:compress.Len()]
	for i, ch := range compress.String() {
		chars[i] = byte(ch)
	}
	return compress.Len()
}

package main

import (
	"slices"
)

func formatLine(line []byte, maxWidth int) string {
	if len(line) > maxWidth {
		return string(line)
	}
	spaces := []int{}
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			spaces = append(spaces, i)
		}
	}
	if len(spaces) == 0 {
		for len(line) < maxWidth {
			line = append(line, ' ')
		}
		return string(line)
	}
Loop:
	for len(line) < maxWidth {
		for i := 0; i < len(spaces); i++ {
			if len(line) >= maxWidth {
				break Loop
			}
			line = slices.Insert(line, spaces[i], ' ')
			for j := i + 1; j < len(spaces); j++ {
				spaces[j]++
			}
		}
	}
	return string(line)
}

func fullJustify(words []string, maxWidth int) []string {
	currentLen := 0
	res := []string{}
	line := []byte{}
	for i := 0; i < len(words); i++ {
		currentLen += len(words[i]) + 1
		if currentLen-1 > maxWidth {
			line = line[:len(line)-1]
			res = append(res, formatLine(line, maxWidth))
			currentLen = len(words[i]) + 1
			line = []byte{}
		}
		line = append(line, []byte(words[i])...)
		line = append(line, ' ')
	}
	line = line[:len(line)-1]
	for len(line) < maxWidth {
		line = append(line, ' ')
	}
	res = append(res, string(line))
	return res
}

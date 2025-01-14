package main

import (
	"regexp"
	"sort"
	"strconv"
)

func answerValidation(first string, second string, third string) string {
	matchedFirst, _ := regexp.MatchString("^(100000|[1-9][0-9]{0,4})$", first)
	if !matchedFirst {
		return "no"
	}
	input := []float64{}
	re := regexp.MustCompile(`[\s]`).Split(second, -1)
	for _, word := range re {
		matched, _ := regexp.MatchString("[0-9]+", word)
		if !matched {
			return "no"
		}
		num, _ := strconv.ParseFloat(word, 64)
		input = append(input, num)
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	re = regexp.MustCompile(`[\s]`).Split(third, -1)
	if len(third) != len(second) {
		return "no"
	}
	for index, word := range re {
		if index > len(input)-1 {
			return "no"
		}
		matched, _ := regexp.MatchString("[0-9]+", word)
		if !matched {
			return "no"
		}
		num, _ := strconv.ParseFloat(word, 64)
		if num != input[index] {
			return "no"
		}
	}
	return "yes"
}

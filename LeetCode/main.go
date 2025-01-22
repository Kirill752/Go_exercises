package main

import "fmt"

func main() {
	// input := []int{4, 2, 0, 3, 2, 5}
	input := "sadbutsad"
	needle := "sad"
	// fmt.Println(input[0:2])
	fmt.Println(strStr(input, needle))
}

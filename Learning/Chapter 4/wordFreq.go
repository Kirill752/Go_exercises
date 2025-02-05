package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordFreq(file *os.File) map[string]int {
	input := bufio.NewScanner(file)
	words := make(map[string]int)
	totalWords := 0
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
		totalWords++
	}
	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "wordFreq: %v\n", input.Err())
	}
	fmt.Printf("Всего слов: %d\n", totalWords)
	return words
}

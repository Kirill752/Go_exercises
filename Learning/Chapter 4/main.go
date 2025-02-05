package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("ReadFile.txt")
	for k, v := range wordFreq(file) {
		fmt.Printf("%s: %d\n", k, v)
	}
}

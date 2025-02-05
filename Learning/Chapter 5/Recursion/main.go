package main

import "fmt"

func main() {
	url := "https://golang.org"
	// n := names(url)
	// for k, v := range n {
	// 	fmt.Printf("%s: %d\n", k, v)
	// }
	words, img, _ := CountWordsAndImages(url)
	fmt.Println(words, img)
}

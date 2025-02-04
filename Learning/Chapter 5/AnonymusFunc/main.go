package main

import (
	"os"
)

func main() {
	// url := "https://golang.org"
	breadthFirst(crawl, os.Args[1:])
}

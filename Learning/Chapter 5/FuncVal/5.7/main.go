package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("Get error: %v", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Parse error: %v", err)
	}
	forEachNode(doc, startElementFull, endElementFull)
}

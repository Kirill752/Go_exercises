package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// names возвращает отображение имя эемента : количество элементов с таким именем
func names(url string) map[string]int {
	countNames := make(map[string]int)
	resp, _ := http.Get(url)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "names: %v\n", err)
		os.Exit(1)
	}
	visitNode(countNames, doc)
	return countNames
}

func visitNode(countNames map[string]int, doc *html.Node) {
	if doc == nil {
		return
	}
	if doc.Type == html.ElementNode {
		countNames[doc.Data]++
	}
	visitNode(countNames, doc.FirstChild)
	visitNode(countNames, doc.NextSibling)
}

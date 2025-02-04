package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func findLinks(url string) {
	resp, _ := http.Get(url)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, doc *html.Node) []string {
	if doc == nil {
		return links
	}
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, name := range doc.Attr {
			if name.Key == "href" {
				links = append(links, name.Val)
			}
		}
	}
	links = visit(links, doc.FirstChild)
	return visit(links, doc.NextSibling)
}

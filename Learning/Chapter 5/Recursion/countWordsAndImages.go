package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndlmages выполняет HTTP-запрос GET HTML-документа
// url и возвращает количество слов и изображений в нем
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	wordsFreq := make(map[string]int)
	words, images = countWordsAndImages(doc, wordsFreq)
	for k, v := range wordsFreq {
		fmt.Printf("%s: %d\n", k, v)
	}
	return
}

func countWordsAndImages(doc *html.Node, wordsFreq map[string]int) (words, images int) {
	if doc == nil {
		return
	}
	if doc.Type == html.ElementNode {
		if doc.Data == "img" {
			images++
		}
	}
	if doc.Type == html.TextNode {
		if len(strings.TrimSpace(doc.Data)) > 0 && doc.Parent.Data != "script" && doc.Parent.Data != "style" {
			w := strings.Fields(doc.Data)
			words += len(w)
			for i := 0; i < len(w); i++ {
				wordsFreq[w[i]]++
			}
		}
	}
	w, i := countWordsAndImages(doc.FirstChild, wordsFreq)
	words += w
	images += i
	w, i = countWordsAndImages(doc.NextSibling, wordsFreq)
	words += w
	images += i
	return
}

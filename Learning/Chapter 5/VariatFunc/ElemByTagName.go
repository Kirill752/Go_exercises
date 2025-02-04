package main

import "golang.org/x/net/html"

// Функция ElementsByTagName возвращает все элементы для данного узла
// дерева HTML, которые соответствуют одному из имен
func ElementsByTagName(doc *html.Node, names ...string) (nodes []*html.Node) {
	if len(names) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, name := range names {
			if doc.Data == name {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, names...)...)
	}
	return
}

package main

import (
	"golang.org/x/net/html"
)

// forEachNode вызывает функции pre(x) и post(x) для каждого узла х
// в дереве с корнем n. Обе функции необязательны.
// рге вызывается до посещения дочерних узлов, a post - после.
// функция прекращает свое выполнение, если функции pre и post возвращают false
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNode(c, id, pre, post); n != nil {
			return n
		}
	}
	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, nil)
}

func findElement(doc *html.Node, id string) bool {
	if doc.Type == html.ElementNode {
		for _, a := range doc.Attr {
			// fmt.Println(a.Key)
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

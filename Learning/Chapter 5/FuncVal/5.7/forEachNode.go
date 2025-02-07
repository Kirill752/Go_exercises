package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// forEachNode вызывает функции pre(x) и post(x) для каждого узла х
// в дереве с корнем n. Обе функции необязательны.
// рге вызывается до посещения дочерних узлов, a post - после,
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func startElementFull(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		if len(n.Attr) == 0 {
			fmt.Printf("<%s/>\n", n.Data)
		} else {
			fmt.Printf("<%s", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s = `%s`", a.Key, a.Val)
			}
			fmt.Print(">")
		}
		depth++
	} else {
		if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
			text := strings.TrimSpace(n.Data)
			if len(text) != 0 {
				fmt.Printf("%s\n", text)
			}
			depth++
		}
	}
}
func endElementFull(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		depth--
		if len(n.Attr) != 0 {
			fmt.Printf("</%s>\n", n.Data)
		}
	}
}

package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type URLReader struct {
	url string
	n   uint64
}

func (r *URLReader) Read(p []byte) (n int, err error) {
	if r.n >= uint64(len(r.url)) {
		return 0, io.EOF
	}
	n = copy(p, r.url[r.n:])
	r.n = uint64(n)
	return
}

func NewReader(s string) io.Reader {
	return &URLReader{url: s, n: 0}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение дескриптора в стек
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

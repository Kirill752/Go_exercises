package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// 7.4
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

// 7.5
type ReaderWithLimit struct {
	R io.Reader
	n int64
}

func (lr *ReaderWithLimit) Read(p []byte) (n int, err error) {
	if lr.n < 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}
	n, err = lr.R.Read(p)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &ReaderWithLimit{R: r, n: n}
}

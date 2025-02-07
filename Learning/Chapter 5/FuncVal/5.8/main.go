package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// go run main.go forEachNode.go -id=client-env -url=https://github.com/
func main() {
	url := flag.String("url", "https://gopl.io", "")
	id := flag.String("id", "toc", "")
	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Printf("Get error: %v", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Parse error: %v", err)
	}
	node := ElementByID(doc, *id)
	if node != nil {
		fmt.Printf("<%s", node.Data)
		for _, a := range node.Attr {
			fmt.Printf(" %s = '%s'", a.Key, a.Val)
		}
		fmt.Print(">\n")
	} else {
		fmt.Println("<nil>")
	}
}

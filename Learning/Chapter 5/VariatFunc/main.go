package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// vals := []int{}
	// maxV, err1 := maxValue(vals...)
	// if err1 != nil {
	// 	log.Println(err1)
	// }
	// minV, err2 := minValue(vals...)
	// if err2 != nil {
	// 	log.Println(err1)
	// }
	// fmt.Println(maxV)
	// fmt.Println(minV)
	// fmt.Println(joinStr(" ", "a", "b", "c", "d", "efg", "h", "ijk", "lmnop", "q", "r", "s"))
	url := "https://golang.org"
	resp, _ := http.Get(url)
	doc, _ := html.Parse(resp.Body)
	resp.Body.Close()
	images := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for i := 0; i < len(images); i++ {
		fmt.Println(images[i])
	}
}

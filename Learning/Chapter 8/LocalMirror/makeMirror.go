package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"

	"gopl.io/ch5/links"
)

var wg sync.WaitGroup
var base = flag.String("b", "https://go.dev", "base cite")

func main() {
	flag.Parse()
	for _, url := range crawl(*base) {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			GetCite(url)
		}(url)
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
	GetCite(*base)
}

func GetCite(url string) {
	m, err := regexp.MatchString(*base, url)
	if err != nil {
		log.Fatal(err)
	}
	if !m {
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatal(fmt.Errorf("getting %s: %s", url, resp.Status))
	}
	defer resp.Body.Close()
	r := regexp.MustCompile(`[^/]+([a-zA-Z0-9-_\.]+)/?$`)
	name := r.FindString(url)
	if name != "" {
		if name[len(name)-1] == '/' {
			name = name[:len(name)-1]
		}
		out, err := os.OpenFile("Links/"+name+"_miror.html", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		out.Close()
	}
}

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества параллельных запросов величиной 20.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

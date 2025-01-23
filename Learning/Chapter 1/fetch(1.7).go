package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// fetch выводит ответ на запрос по заданному URL
// Пример запуска:
// ./execute http://gopl.io
func fetch() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Код состояния HTTP
		fmt.Printf("HTTP status: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

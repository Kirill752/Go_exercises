package main

import (
	"log"
	"net/http"
)

func main() {
	// fetch()
	http.HandleFunc("/", url_path) // Каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", handler_lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

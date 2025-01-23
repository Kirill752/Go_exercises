package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerPlot3D)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

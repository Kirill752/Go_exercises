package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerPlot3D)
	http.HandleFunc("/MandelBrot", handleMandelbrot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

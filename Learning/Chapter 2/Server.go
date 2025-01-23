package main

import (
	"math"
	"net/http"
	"strconv"
)

func handlerPlot3D(w http.ResponseWriter, r *http.Request) {
	// Задаем параметры отображения
	setParams(r)
	// Задаем параметры отображения ответа с сервера
	w.Header().Set("Content-Type", "image/svg+xml")
	plot3D(w)
}

// Задание параметров в http запросе
// Пример:
// http://localhost:8000/count?w=1000&h=1000
func setParams(r *http.Request) {
	if canvasWidth, err := strconv.ParseFloat(r.URL.Query().Get("w"), 64); err == nil {
		width = canvasWidth
	}
	if canvasHeight, err := strconv.ParseFloat(r.URL.Query().Get("h"), 64); err == nil {
		height = canvasHeight
	}
	if ang, err := strconv.ParseFloat(r.URL.Query().Get("angle"), 64); err == nil {
		angle = ang * math.Pi / 180
	}
	strk := r.URL.Query().Get("stroke")
	if strk != "" {
		stroke = strk
	}
	fll := r.URL.Query().Get("fill")
	if fll != "" {
		fill = fll
	}
}

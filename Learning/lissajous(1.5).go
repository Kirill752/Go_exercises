package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

// Программа формирует .gif картинку фигур Лиссажуэ
// Пример запуска:
// ./execute >lissajous.gif
func lissajous(out io.Writer) {
	pallete := []color.Color{color.Black, color.White, color.RGBA{0, 255, 0, 255}, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}}
	const (
		cycles  = 5     // количество полных колебаний x
		res     = 0.001 // Угловое разришение
		size    = 100   // Канва изображения
		nframes = 64    // Количество кадров анимации
		delay   = 1     // задержка между кадрами (1 - 10 мс)
	)
	const (
		blackIndex = 0
		whiteIndex = 1
		greenIndex = 2
		redIndex   = 3
		blueIndex  = 4
	)
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := rand.Float64() * 3 // относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += 0.01 {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size), size+int(y*size), uint8((rand.Int())%len(pallete))) // третий индекс тут обозначает индекс цвета в палитре pallete
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

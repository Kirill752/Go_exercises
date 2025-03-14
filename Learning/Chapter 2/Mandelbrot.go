package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func mandelbrot(z complex128) color.Color {
	iterations := 200
	contrast := 15
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	// return color.YCbCr{1, 5, 8}
	return color.Transparent
}

func printMandelbrot(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // Примечание: игнорируем ошибки
}

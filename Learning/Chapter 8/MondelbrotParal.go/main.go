package main

import (
	"fmt"
	"image/png"
	"os"
	"time"
)

func main() {
	out, _ := os.OpenFile("Mondelbrot.png", os.O_WRONLY, os.ModePerm)
	defer out.Close()

	t := time.Now()
	img := parallelRender(7)
	png.Encode(out, img)
	mst := time.Since(t)
	fmt.Printf("mandelbrot: paralel: %d ms\n", mst.Milliseconds())

	t = time.Now()
	serialRender(out)
	mst = time.Since(t)
	fmt.Printf("mandelbrot: serial: %d ms\n", mst.Milliseconds())
}

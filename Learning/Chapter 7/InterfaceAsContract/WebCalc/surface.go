package main

import (
	"fmt"
	"io"
	"math"
)

// Plot3D вычисляет SVG-представление трехмерного графика функции.

var width, height = 600.0, 320.0         // Размер канвы в пикселях
var cells = 100.0                        // Количество ячеек сетки
var xyrange = 30.0                       // Диапазон осей (-xyrange, ..., xyrange)
var xyscale = float64(width/2) / xyrange // Пикселей в еденице  x или y
var zscale = 0.4 * height                // Пикселей в еденице z
var angle = math.Pi / 6                  //Угол осей x, y
var stroke = "gray"
var fill = "white"

// Вычисляет значения (x, y, z) функции z = f(x, y)
// i, j - номер ячейки на двумерной сетке размером cells*cells
// исходя из диапазона значений xyrange.
// Затем проецирует эти значения на двумерную канву
func corner(i, j int, f func(float64, float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумерную канву (sx, sy)
	sx := width/2 + (x-y)*math.Cos(angle)*xyscale
	sy := height/2 + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy
}

// Функция для построения
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние до (0, 0)
	return math.Sin(r) / r
}
func plot3D(out io.Writer, fn func(float64, float64) float64) {
	// output, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%.f' height='%.f'>", stroke, fill, width, height)
	// output.Write([]byte(label))
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(i+1, j, fn)
			bx, by := corner(i, j, fn)
			cx, cy := corner(i, j+1, fn)
			dx, dy := corner(i+1, j+1, fn)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %gj%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			// output.Write([]byte(text))
		}
	}
	fmt.Fprint(out, "</svg>")
	// output.Write([]byte("</svg>"))
	// output.Close()
}

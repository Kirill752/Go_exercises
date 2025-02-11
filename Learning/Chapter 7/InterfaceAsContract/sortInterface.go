package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type less func(x, y *Track) bool

func colTitle(x, y *Track) bool  { return x.Title < y.Title }
func colArtist(x, y *Track) bool { return x.Artist < y.Artist }
func colAlbum(x, y *Track) bool  { return x.Album < y.Album }
func colYear(x, y *Track) bool   { return x.Year < y.Year }
func colLength(x, y *Track) bool { return x.Length < y.Length }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "---- ", "------", "-----", "--- ", "----- ")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist,
			t.Album, t.Year, t.Length)
	}
	tw.Flush() // Вычисление размеров столбцов и вывод таблицы
}
func printTracksHTML(tracks []*Track) {
	var report = template.Must(template.ParseFiles("./temp.html"))
	out, err := os.OpenFile("Tracks.html", os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	err = report.Execute(out, tracks)
	if err != nil {
		log.Println(err)
	}
}
func printTracksFile(out io.Writer, tracks []*Track) {
	var report = template.Must(template.ParseFiles("./temp.html"))
	err := report.Execute(out, tracks)
	if err != nil {
		log.Println(err)
	}
}

// Сортировка по названию
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Сортировка по Артисту
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Сортировка по Альбому
type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Сортировка по году
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Сортировка по времени
type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Кастомная сортировка
type customSort struct {
	t    []*Track
	less func(x, у *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

type userSort struct {
	t            []*Track
	sortPriority []less
}

func (x userSort) Len() int { return len(x.t) }
func (x userSort) Less(i, j int) bool {
	I, J := x.t[i], x.t[j]
	for _, f := range x.sortPriority {
		switch {
		case f(I, J):
			return true
		case f(J, I):
			return false
		}
	}
	// Если все столбцы оказались одинаковыми, то используем первый столбец
	return x.sortPriority[0](I, J)
}
func (x userSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func sortbyCol(t []*Track, F ...less) *userSort {
	usrSrt := &userSort{t: t, sortPriority: F}
	if sort.IsSorted(usrSrt) {
		sort.Sort(sort.Reverse(usrSrt))
	} else {
		sort.Sort(usrSrt)
	}
	return usrSrt
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/title":
		sortbyCol(tracks, colTitle)
	case "/artist":
		sortbyCol(tracks, colArtist)
	case "/album":
		sortbyCol(tracks, colAlbum)
	case "/year":
		sortbyCol(tracks, colYear)
	case "/length":
		sortbyCol(tracks, colLength)
	}
	printTracksFile(w, tracks)
}

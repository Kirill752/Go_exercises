package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	// var fl os.File
	w = os.Stdout
	f := w.(*os.File)
	c := w.(*os.File)
	fmt.Printf("f (%T)\n", f)
	fmt.Printf("c (%T)\n", c)
}

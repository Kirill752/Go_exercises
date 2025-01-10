package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var number string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &number)
		fmt.Fprintf(out, "%s\n", deleteNumber(number))
	}
}

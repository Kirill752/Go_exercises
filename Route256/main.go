package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t string
	t, _ = in.ReadString('\n')
	n, _ := strconv.Atoi(t[:len(t)-1])
	var first, second, third string
	for i := 0; i < n; i++ {
		first, _ = in.ReadString('\n')
		second, _ = in.ReadString('\n')
		third, _ = in.ReadString('\n')
		fmt.Fprintf(out, "%s\n", answerValidation(first[:len(first)-1], second[:len(second)-1], third[:len(third)-1]))
	}
}

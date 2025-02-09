package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stdout, "wordCounter:", err)
	}
	return int(*c), err
}
func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stdout, "lineCounter:", err)
	}
	return int(*c), err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newW := bufio.NewWriter(w)
	sz := int64(newW.Size())
	return newW, &sz
}

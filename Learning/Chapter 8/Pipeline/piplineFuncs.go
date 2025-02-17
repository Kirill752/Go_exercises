package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}
func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Printf("%v ", x)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

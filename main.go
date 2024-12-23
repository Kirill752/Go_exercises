package main

import (
	"fmt"
)

func main() {
	i, j := 1, 270

	p := &i
	fmt.Printf("The type of i = %d is %T \n", i, i)
	fmt.Printf("The type of j = %d is %T \n", j, i)
	fmt.Printf("The adress of i = %d is %d \n", i, p)
	fmt.Printf("The adress of j = %d is %d \n", j, &j)
	fmt.Println(p)
	fmt.Println(&j)	
}

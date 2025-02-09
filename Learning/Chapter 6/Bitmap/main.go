package main

import "fmt"

func main() {
	var x IntSet
	x.Add(1)
	x.Add(4)
	x.Remove(1)
	x.Add(52)
	y := x.Copy()
	y.AddAll(6, 0, 52, 145)
	x.AddAll(10, 33, 145)
	fmt.Printf("x = %v\n", &x)
	fmt.Printf("y = %v\n", y)

	y.SymetricDifferenceWith(&x)
	fmt.Printf("y SymetricDifferenceWith x = %v\n", y)
	fmt.Println(x.Elems())
}

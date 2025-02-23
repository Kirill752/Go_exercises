package main

import "fmt"

func main() {
	equations := [][]string{{"a", "b"}, {"a", "c"}, {"d", "e"}, {"d", "f"}, {"a", "d"}, {"aa", "bb"}, {"aa", "cc"}, {"dd", "ee"}, {"dd", "ff"}, {"aa", "dd"}, {"a", "aa"}}
	values := []float64{2.0, 3.0, 4.0, 5.0, 7.0, 5.0, 8.0, 9.0, 3.0, 2.0, 2.0}
	queries := [][]string{{"ff", "a"}}
	ans := calcEquation(equations, values, queries)
	fmt.Println(ans)
}

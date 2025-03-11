package main

func Calc(Q []int, C []int, A, B int) int {
	n := len(Q)
	DQ := 0
	if A == B {
		for i := range n {
			DQ += (A * Q[i])
		}
		return DQ
	}
	for i := range n {
		DQ += ((C[i]*(B-A))/255 + A) * Q[i]
	}
	return DQ
}

// func main() {
// 	Q := []int{1000, 2000, 3000, 4000, 5000}
// 	C := []int{0, 63, 127, 191, 255}
// 	A, B := 1000, 5000
// 	fmt.Println(Calc(Q, C, A, B))
// }

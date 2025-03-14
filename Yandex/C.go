package main

type TAXI struct {
	timestamp int
	id        int
	position  int
}

type ORDER struct {
	timestamp int
	id        int
	position  int
	t         int
}

func canReach(taxies []TAXI, ord ORDER, L, S int) []int {
	maxDist := ord.t * S
	res := []int{}
	// Граница области с которой могут доехать таксисты
	leftBoarder := ord.position - maxDist
	if leftBoarder < 0 {
		leftBoarder = L - (maxDist - ord.position)
	}
	for i := range taxies {
		rightBoarder := ord.position - (ord.timestamp-taxies[i].timestamp)*S
		if rightBoarder < 0 {
			rightBoarder = L - ((ord.timestamp-taxies[i].timestamp)*S - ord.position)
		}
		if taxies[i].position >= leftBoarder && taxies[i].position <= rightBoarder {
			res = append(res, taxies[i].id)
		}
	}
	return res
}

// func main() {
// 	t1 := TAXI{1738300000, 0, 0}
// 	t2 := TAXI{1738300000, 1, 2}
// 	o1 := ORDER{1738300001, 0, 1, 1}
// 	fmt.Println(canReach([]TAXI{t1, t2}, o1, 100, 1))
// }

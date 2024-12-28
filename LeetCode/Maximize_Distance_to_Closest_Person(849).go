package main

func maxDistToClosest(seats []int) int {
	// Расстояние до левых соседей
	l_r := []int{}
	// Расстояние до правых соседей
	r_l := []int{}
	len_to_neighbour := 0
	// Маркер того, что мы встретили хотябы одного человека
	marker := false
	// Обход слева направо
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			marker = true
			len_to_neighbour = 0
			l_r = append(l_r, len_to_neighbour)
		} else {
			if marker == true {
				len_to_neighbour++
				l_r = append(l_r, len_to_neighbour)
			} else {
				l_r = append(l_r, len(seats)+1)
			}
		}
	}
	// обход справа налево
	marker = false
	len_to_neighbour = 0
	for i := len(seats) - 1; i > -1; i-- {
		if seats[i] == 1 {
			marker = true
			len_to_neighbour = 0
			r_l = append(r_l, len_to_neighbour)
		} else {
			if marker == true {
				len_to_neighbour++
				r_l = append(r_l, len_to_neighbour)
			} else {
				r_l = append(r_l, len(seats)+1)
			}
		}
	}

	result := 0
	buf := 0
	size := len(l_r) - 1
	for i := 0; i <= size; i++ {
		buf = min(l_r[i], r_l[size-i])
		if buf > result {
			result = buf
		}
	}
	return result
}

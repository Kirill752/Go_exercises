package main

import "fmt"

func ASCIIRobots(field [][]byte) [][]byte {
	coordsA := make([]int, 2)
	coordsB := make([]int, 2)
	findA, findB := false, false
	// Ищем место нахождение роботов A и B
	for i, str := range field {
		for j, ch := range str {
			if ch == 'A' {
				coordsA[0] = i
				coordsA[1] = j
			}
			if ch == 'B' {
				coordsB[0] = i
				coordsB[1] = j
			}
			if findA && findB {
				break
			}
		}
	}
	fmt.Printf("A = %v\n", coordsA)
	fmt.Printf("B = %v\n", coordsB)
	return field
}

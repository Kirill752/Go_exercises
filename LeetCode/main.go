package main

import "fmt"

func main() {
	startGene := "hit"
	endGene := "cog"
	bank := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(startGene, endGene, bank))
}

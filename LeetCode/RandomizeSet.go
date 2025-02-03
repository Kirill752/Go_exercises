package main

import (
	"math/rand"
)

type RandomizedSet struct {
	setMap map[int]bool
	keys   []int
}

func ConstructorSet() RandomizedSet {
	example := RandomizedSet{}
	example.setMap = make(map[int]bool)
	return example
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.setMap[val]; ok {
		return false
	} else {
		this.setMap[val] = true
		this.keys = append(this.keys, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.setMap[val]; ok {
		delete(this.setMap, val)
		for i := 0; i < len(this.keys); i++ {
			if this.keys[i] == val {
				copy(this.keys[i:], this.keys[i+1:])
				this.keys = this.keys[:len(this.keys)-1]
			}
		}
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.keys))
	return this.keys[index]
}

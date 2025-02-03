package main

import (
	"container/list"
)

type LRUCache struct {
	// key-value
	data map[int]int
	// key-listelem
	age map[int]*list.Element
	// capacity
	capacity int
	lst      *list.List
}

func Constructor(capacity int) LRUCache {
	var newCache LRUCache
	newCache.data = make(map[int]int)
	newCache.age = make(map[int]*list.Element)
	newCache.lst = list.New()
	newCache.capacity = capacity
	return newCache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		this.lst.MoveToBack(this.age[key])
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.data[key] = value
		this.lst.MoveToBack(this.age[key])
	} else {
		this.data[key] = value
		// fmt.Println(this.lst.Len(), this.capacity)
		if this.lst.Len() >= this.capacity {
			delete(this.data, this.lst.Front().Value.(int))
			delete(this.age, this.lst.Front().Value.(int))
			this.lst.Remove(this.lst.Front())
		}
		this.lst.PushBack(key)
		this.age[key] = this.lst.Back()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

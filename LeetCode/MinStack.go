package main

import "slices"

type MinStack struct {
	stack []int
	minEl int
}

func ConstructorS() MinStack {
	newstack := MinStack{stack: []int{}}
	return newstack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack) == 1 {
		this.minEl = val
	} else {
		this.minEl = min(this.minEl, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 1 {
		if this.stack[len(this.stack)-1] == this.minEl {
			this.stack = this.stack[:len(this.stack)-1]
			this.minEl = slices.Min(this.stack)
		} else {
			this.stack = this.stack[:len(this.stack)-1]
		}
	} else {
		if len(this.stack) > 0 {
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minEl
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l6 := new(ListNode)
	l7 := new(ListNode)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7

	// l1.Random = nil
	// l2.Random = l1
	// l3.Random = l5
	// l4.Random = l3
	// l5.Random = l1

	l1.Val = 1
	l2.Val = 4
	l3.Val = 3
	l4.Val = 0
	l5.Val = 2
	l6.Val = 5
	l7.Val = 2

	h := l1
	i := 1
	h = partition(l1, 3)
	// fmt.Println()
	for ; h != nil; h = h.Next {
		fmt.Printf("%d: (%d, %v)\n", i, h.Val, h.Next)
		i++
	}
	// fmt.Println()
}

package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = nil

	// l1.Random = nil
	// l2.Random = l1
	// l3.Random = l5
	// l4.Random = l3
	// l5.Random = l1

	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 4
	l5.Val = 5

	h := l1
	i := 1
	h = reverseKGroup(l1, 2)
	fmt.Println()
	for ; h != nil; h = h.Next {
		fmt.Printf("%d: (%d, %v)\n", i, h.Val, h.Next)
		i++
	}
	// fmt.Println()
}

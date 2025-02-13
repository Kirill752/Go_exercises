package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l1.Next = new(ListNode)
	l1.Next.Next = new(ListNode)
	l1.Val = 1
	l1.Next.Val = 2
	l1.Next.Next.Val = 4
	l2 := new(ListNode)
	l2.Next = new(ListNode)
	l2.Next.Next = new(ListNode)
	l2.Val = 1
	l2.Next.Val = 3
	l2.Next.Next.Val = 4
	r := mergeTwoLists(l1, l2)
	for ; r != nil; r = r.Next {
		fmt.Printf("%d ", r.Val)
	}
	fmt.Println()
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var follow *ListNode
	for n := head; n != nil; {
		follow = n.Next
		n.Next = prev
		prev = n
		n = follow
	}
	return prev
}

func reversePart(start *ListNode, end *ListNode) *ListNode {
	var prev *ListNode
	var follow *ListNode
	stop := end.Next
	for n := start; n != stop; {
		follow = n.Next
		n.Next = prev
		prev = n
		n = follow
	}
	return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	left--
	right--
	n := head
	var prevLeft *ListNode
	for ; left > 0 && n != nil; left-- {
		prevLeft = n
		n = n.Next
	}
	leftNode := n
	n = head
	for ; right > 0 && n != nil; right-- {
		n = n.Next
	}
	rightNode := n
	folowRight := rightNode.Next
	r := reversePart(leftNode, rightNode)
	if prevLeft != nil {
		prevLeft.Next = r
		leftNode.Next = folowRight
		return head
	} else {
		leftNode.Next = folowRight
		return r
	}
}

func reverseKGroup(head *ListNode, k int) (res *ListNode) {
	L := 0
	stop := 0
	res = head
	for n := head; n != nil; n = n.Next {
		stop++
	}
	for L < stop {
		R := L + k - 1
		if R >= stop {
			break
		}
		res = reverseBetween(res, L+1, R+1)
		L = R + 1
	}
	return
}

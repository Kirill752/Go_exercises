package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	n := new(Node)
	Head := n
	mapAdress := map[*Node]*Node{}
	for N := head; N != nil; N = N.Next {
		mapAdress[N] = n
		n.Val = N.Val
		n.Next = N.Next
		n.Random = N.Random
		newNode := new(Node)
		n = newNode
	}
	for N := Head; N != nil; N = N.Next {
		N.Next = mapAdress[N.Next]
		N.Random = mapAdress[N.Random]
	}
	return Head
}

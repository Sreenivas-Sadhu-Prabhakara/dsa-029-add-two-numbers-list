package main

import "fmt"

// Node is a singly linked list node holding one decimal digit.
type Node struct {
	Val  int
	Next *Node
}

// AddTwoNumbers sums two non-negative integers stored as linked lists with
// the least significant digit first, returning the sum in the same format.
func AddTwoNumbers(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	carry := 0
	for a != nil || b != nil || carry != 0 {
		sum := carry
		if a != nil {
			sum += a.Val
			a = a.Next
		}
		if b != nil {
			sum += b.Val
			b = b.Next
		}
		tail.Next = &Node{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}
	return dummy.Next
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	// 342 + 465 = 807, stored least significant digit first.
	fmt.Println(toSlice(AddTwoNumbers(fromSlice([]int{2, 4, 3}), fromSlice([]int{5, 6, 4}))))
}

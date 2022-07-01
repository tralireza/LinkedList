package LinkedList

import "log"

func init() {
	log.SetFlags(0)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2816m Double a Number Represented as a Linked List
func doubleIt(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := doubleIt(head.Next)

	head.Val *= 2
	if n != head.Next {
		head.Val += 1
	}
	if head.Val >= 10 {
		head.Val %= 10
		return &ListNode{1, head}
	}
	return head
}

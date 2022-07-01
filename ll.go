package LinkedList

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (o *ListNode) String() string {
	if o.Next != nil {
		return fmt.Sprintf("{%d *}->", o.Val)
	}
	return fmt.Sprintf("{%d /}", o.Val)
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

package LinkedList

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func init() {}

// 2816m Double a Number Represented as a Linked List
func Test2816(t *testing.T) {
	Draw := func(head *ListNode) string {
		var bfr strings.Builder
		for n := head; n != nil; n = n.Next {
			if n.Next != nil {
				bfr.WriteString(fmt.Sprintf("{%d *}->", n.Val))
			} else {
				bfr.WriteString(fmt.Sprintf("{%d /}", n.Val))
			}
		}
		return bfr.String()
	}

	Draw2 := func(head *ListNode) string {
		var bfr bytes.Buffer
		for n := head; n != nil; n = n.Next {
			fmt.Fprintf(&bfr, "%v", n)
		}
		return bfr.String()
	}

	WithReverse := func(head *ListNode) *ListNode {
		Reverse := func(head *ListNode) *ListNode {
			var prv, nxt *ListNode
			for n := head; n != nil; n = nxt {
				prv, n.Next, nxt = n, prv, n.Next
			}
			return prv
		}

		head = Reverse(head)

		var carry int
		var prv *ListNode
		for n := head; n != nil; n = n.Next {
			n.Val *= 2
			n.Val += carry
			if n.Val >= 10 {
				n.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
			prv = n
		}
		if carry == 1 {
			prv.Next = &ListNode{Val: 1}
		}

		return Reverse(head)
	}

	type L = ListNode
	for _, f := range []func(*ListNode) *ListNode{WithReverse, doubleIt} {
		l := &L{9, &L{9, &L{Val: 9}}}
		log.Printf("%v -> %v", Draw(l), Draw2(f(l)))
	}
}

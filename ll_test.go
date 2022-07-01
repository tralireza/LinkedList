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

	type L = ListNode
	for _, l := range []*ListNode{&L{9, &L{9, &L{Val: 9}}}, &L{1, &L{2, &L{Val: 3}}}} {
		log.Printf("%v -> %v", Draw(l), Draw2(doubleIt(l)))
	}
}

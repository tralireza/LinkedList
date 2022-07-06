package LinkedList

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func init() {}

// 24m Swap Nodes in Pairs
func Test24(t *testing.T) {
	Draw := func(head *ListNode) string {
		var bfr bytes.Buffer
		for n := head; n != nil; n = n.Next {
			fmt.Fprintf(&bfr, "%v", n)
		}
		return bfr.String()
	}

	type L = ListNode

	log.Print(" -> ", Draw(swapPairs(&L{1, &L{2, &L{3, &L{Val: 4}}}})))
	log.Print(" -> ", Draw(swapPairs(&L{Val: 1})))
}

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

	TwoPointer := func(head *ListNode) *ListNode {
		var prv *ListNode
		for n := head; n != nil; n = n.Next {
			n.Val *= 2
			if n.Val >= 10 {
				n.Val %= 10
				if prv != nil {
					prv.Val++
				} else {
					head = &ListNode{1, head}
				}
			}
			prv = n
		}

		return head
	}

	type L = ListNode
	for _, f := range []func(*ListNode) *ListNode{TwoPointer, WithReverse, doubleIt} {
		log.Print("== ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		for _, l := range []*ListNode{&L{9, &L{9, &L{Val: 9}}}, &L{1, &L{Val: 5}}} {
			log.Printf("%v -> %v", Draw(l), Draw2(f(l)))
		}
	}
}

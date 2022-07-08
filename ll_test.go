package LinkedList

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strconv"
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

// 138m Copy List with Random Pointer
func Test138(t *testing.T) {
	type Node struct {
		Val    int
		Next   *Node
		Random *Node
	}

	copyRandomList := func(head *Node) *Node {
		Mem := map[*Node]*Node{}

		var cpyhead, prv *Node
		for n := head; n != nil; n = n.Next {
			cpy := &Node{Val: n.Val * 10}
			Mem[n] = cpy

			if cpyhead == nil {
				cpyhead = cpy
			} else {
				prv.Next = cpy
			}
			prv = cpy
		}

		cpy := cpyhead
		for n := head; n != nil; n = n.Next {
			if n.Random != nil {
				cpy.Random = Mem[n.Random]
			}
			cpy = cpy.Next
		}

		return cpyhead
	}

	Draw := func(head *Node) {
		for n := head; n != nil; n = n.Next {
			x := "*}->"
			if n.Next == nil {
				x = "/}"
			}
			r := "/"
			if n.Random != nil {
				r = "[" + strconv.Itoa(n.Random.Val) + "]"
			}
			fmt.Printf("{%d %s %s", n.Val, r, x)
		}
		fmt.Print("\n")
	}

	type N = Node
	n1, n2, n3, n4, n5 := &N{Val: 1}, &N{Val: 2}, &N{Val: 3}, &N{Val: 4}, &N{Val: 5}
	n1.Next, n2.Next, n3.Next, n4.Next = n2, n3, n4, n5
	n1.Random, n3.Random, n5.Random = n3, n1, n2

	Draw(copyRandomList(&N{0, n1, n3}))
}

// 142m Linked List Cycle II
func Test142(t *testing.T) {
	type L = ListNode

	l4 := &L{Val: -4}
	l := &L{3, &L{2, &L{0, l4}}}
	l4.Next = l
	log.Print(" ?= ", detectCycle(l))

	l4.Next = l.Next
	log.Print(" ?= ", detectCycle(l))

	log.Print(" ?= ", detectCycle(&L{Val: 1}))
	log.Print(" ?= ", detectCycle(&L{1, &L{Val: 2}}))
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

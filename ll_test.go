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

// 2m Add Two Numbers
func Test2(t *testing.T) {
	Draw := func(head *ListNode) string {
		var bfr strings.Builder
		for n := head; n != nil; n = n.Next {
			bfr.WriteString(n.String())
		}
		return bfr.String()
	}

	type L = ListNode
	for _, l := range [][]*L{
		[]*L{&L{9, &L{9, &L{9, &L{9, &L{9, &L{9, &L{Val: 9}}}}}}}, &L{9, &L{9, &L{9, &L{Val: 9}}}}},
		[]*L{&L{2, &L{4, &L{Val: 3}}}, &L{5, &L{6, &L{Val: 4}}}},
		[]*L{&L{2, &L{Val: 3}}, &L{5, &L{Val: 7}}},
		[]*L{&L{Val: 9}, &L{1, &L{Val: 9}}},
	} {
		log.Printf("%v  |+|  %v  ==  %v", Draw(l[0]), Draw(l[1]), Draw(addTwoNumbers(l[0], l[1])))
	}
}

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
	WithMap := func(head *ListNode) *ListNode {
		M := map[*ListNode]struct{}{}
		for n := head; n != nil; n = n.Next {
			if _, ok := M[n]; ok {
				return n
			}
			M[n] = struct{}{}
		}
		return nil
	}

	type L = ListNode

	for _, f := range []func(*ListNode) *ListNode{detectCycle, WithMap} {
		log.Print("==")

		l4 := &L{Val: -4}
		l := &L{3, &L{2, &L{1, &L{0, l4}}}}
		l4.Next = l
		log.Print(" ?= ", f(l))

		l4.Next = l.Next
		log.Print(" ?= ", f(l))

		log.Print(" ?= ", f(&L{Val: 1}))
		log.Print(" ?= ", f(&L{1, &L{Val: 2}}))
	}
}

// 146m LRU Cache
func Test146(t *testing.T) {
	lru := NewLRUCache(3)
	for _, n := range []int{1, 2, 3, 4, 3, 1, 5, 1} {
		lru.Put(n, n)
		lru.Draw()
	}
	for _, n := range lru.kMap {
		log.Printf("%p -> %[1]v", n)
	}

	for _, n := range []int{1, 2, 3, 3, 1, 2, 4, 5, 3} {
		if lru.Get(n) == -1 {
			log.Printf("%d -> -1", n)
		}
		lru.Draw()
	}
}

// 148m Sort List
func Test148(t *testing.T) {
	Draw := func(head *ListNode) string {
		var bfr bytes.Buffer
		for n := head; n != nil; n = n.Next {
			if n.Next != nil {
				fmt.Fprintf(&bfr, "{%d *}->", n.Val)
			} else {
				fmt.Fprintf(&bfr, "{%d /}", n.Val)
			}
		}
		return bfr.String()
	}

	MergeSort := func(head *ListNode) *ListNode {
		Merge := func(l, r *ListNode) *ListNode {
			s := &ListNode{} // Head Sentinel

			n := s
			for l != nil && r != nil {
				if l.Val <= r.Val {
					n.Next, l = l, l.Next
				} else {
					n.Next, r = r, r.Next
				}
				n = n.Next
			}

			if l != nil {
				n.Next = l
			}
			if r != nil {
				n.Next = r
			}

			return s.Next
		}

		var Sort func(*ListNode) *ListNode
		Sort = func(n *ListNode) *ListNode {
			if n == nil || n.Next == nil {
				return n
			}

			var prv *ListNode
			rN, f := n, n
			for f != nil && f.Next != nil {
				prv = rN
				rN, f = rN.Next, f.Next.Next
			}
			prv.Next = nil

			return Merge(Sort(n), Sort(rN))
		}

		return Sort(head)
	}

	type L = ListNode
	for _, f := range []func(*L) *L{sortList, MergeSort} {
		log.Print("==", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		for _, l := range []*L{&L{4, &L{2, &L{1, &L{Val: 3}}}}, &L{-1, &L{5, &L{3, &L{4, &L{Val: 0}}}}}, &L{3, &L{Val: 1}}} {
			log.Print(Draw(l), "  ->  ", Draw(f(l)))
		}
	}
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

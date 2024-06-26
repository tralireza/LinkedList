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

// 2m Add Two Numbers
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var lSum func(l1, l2 *ListNode, carry int) *ListNode
	lSum = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil {
			if carry > 0 {
				return &ListNode{Val: carry}
			}
			return nil
		}

		if l1 == nil {
			return lSum(&ListNode{Val: carry}, l2, 0)
		}
		if l2 == nil {
			return lSum(&ListNode{Val: carry}, l1, 0)
		}

		n := &ListNode{Val: l1.Val + l2.Val + carry}
		carry = n.Val / 10
		n.Val %= 10
		n.Next = lSum(l1.Next, l2.Next, carry)
		return n
	}

	return lSum(l1, l2, 0)
}

// 24m Swap Nodes in Pairs
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := swapPairs(head.Next.Next)
	head, head.Next.Next, head.Next = head.Next, head, r
	return head
}

// 142m Linked List Cycle II
func detectCycle(head *ListNode) *ListNode {
	p, p2 := head, head
	for p2 != nil && p2.Next != nil {
		log.Print(p, p2)
		p = p.Next
		p2 = p2.Next.Next

		if p == p2 {
			log.Print(" -> LOOPING :: ", p)
			link := head
			for n := p.Next; ; n = n.Next {
				if n == link {
					return link
				}
				if n == p {
					link = link.Next
				}
			}
		}
	}
	log.Print(" -> NONE")
	return nil
}

// 146m LRU Cache
type LRUCache struct {
	kMap       map[int]*dblNode
	Head, Tail *dblNode
	Size, Cap  int
}

type dblNode struct {
	Key, Val   int
	Prev, Next *dblNode
}

func (o dblNode) String() string {
	return fmt.Sprintf("{<%p [%d]:%d %p>}", o.Prev, o.Key, o.Val, o.Next)
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		kMap: map[int]*dblNode{},
		Cap:  capacity,
	}
}

func (lru *LRUCache) Draw() {
	if lru.Size == 0 {
		fmt.Printf("#%d:%d {} \n", lru.Size, lru.Cap)
		return
	}
	fmt.Printf("#%d:%d :: { %d:%d } ||  ", lru.Size, lru.Cap, lru.Head.Key, lru.Head.Val)
	var prv *dblNode
	for n := lru.Head; prv != lru.Tail; n = n.Next {
		fmt.Printf("{- %d:%d -}", n.Key, n.Val)
		prv = n
	}
	fmt.Printf("  || { %d:%d }   -> ", lru.Tail.Key, lru.Tail.Val)

	n := lru.Head
	for range lru.Size {
		fmt.Printf("{%d:%d}", n.Key, n.Val)
		n = n.Next
	}
	fmt.Print("\n")
}

func (lru *LRUCache) Get(key int) int {
	if n, ok := lru.kMap[key]; ok {
		switch n {
		case lru.Head:
			lru.Head, lru.Tail = lru.Head.Next, lru.Tail.Next // rotate: 1 shift to next: <n> moves to tail

		case lru.Tail:

		default:
			n.Prev.Next, n.Next.Prev = n.Next, n.Prev // splice out
			n.Next, n.Prev = lru.Head, lru.Tail       // move to Tail
			n.Prev.Next, n.Next.Prev = n, n           // connect back in
			lru.Tail = n
		}
		return n.Val
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if n, ok := lru.kMap[key]; ok {
		n.Val = value
		lru.Get(key)
		return
	}

	n := &dblNode{Key: key, Val: value}

	switch lru.Size {
	case 0:
		lru.Head, lru.Tail, n.Next, n.Prev = n, n, n, n

	case lru.Cap:
		delete(lru.kMap, lru.Head.Key) // Head is evicted

		lru.Size--
		if lru.Size == 0 {
			lru.Put(key, value)
			return
		}

		// Head is gone, wire up the new Head in
		lru.Head = lru.Head.Next
		lru.Head.Prev = lru.Tail
		lru.Tail.Next = lru.Head

		lru.Put(n.Key, n.Val)
		return

	default:
		n.Next, n.Prev = lru.Head, lru.Tail
		n.Prev.Next, n.Next.Prev = n, n
		lru.Tail = n
	}

	lru.kMap[n.Key] = n
	lru.Size++
}

// 148m Sort List
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tSort := sortList(head.Next)

	if head.Val <= tSort.Val {
		head.Next = tSort
		return head
	}

	prv := head
	for n := tSort; n != nil && head.Val > n.Val; n = n.Next {
		prv = n
	}
	prv.Next, head.Next = head, prv.Next
	return tSort
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

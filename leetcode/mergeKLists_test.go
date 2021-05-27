package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	if k == 1 {
		return lists[0]
	}

	prev := lists[0]
	for i := 1; i < k; i++ {
		prev = mergeTwoLists(prev, lists[i])
	}
	return prev
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var root *ListNode

	if l1.Val < l2.Val {
		root = l1
		l1 = l1.Next
	} else {
		root = l2
		l2 = l2.Next
	}

	p := root
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return root
}

func TestMergeTwoLists(t *testing.T) {
	l1 := buildListNode([]int{1, 4, 5})
	l2 := buildListNode([]int{1, 2, 3, 4, 6})
	printListNode(l1)
	printListNode(l2)

	l3 := mergeTwoLists(l1, l2)
	printListNode(l3)
}

func BenchmarkMergeTwoLists(b *testing.B) {

	for i := 0; i < b.N; i++ {
		l1 := buildListNode([]int{1, 4, 5})
		l2 := buildListNode([]int{1, 2, 3, 4, 6})
		mergeTwoLists(l1, l2)
	}
}

func buildListNode(lst []int) *ListNode {
	root := &ListNode{Val: lst[0]}
	p := root

	for i := 1; i < len(lst); i++ {
		nd := &ListNode{Val: lst[i]}
		p.Next = nd
		p = nd
	}

	return root
}

func printListNode(root *ListNode) {
	p := root
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

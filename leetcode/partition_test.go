package main

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	var lt, gt, ltHead, gtHead *ListNode

	curr := head

	for curr != nil {

		if curr.Val < x {
			if lt != nil {
				lt.Next = curr
				lt = curr
			} else {
				ltHead = curr // head of less numbers
				lt = curr
			}
		} else {
			if gt != nil {
				gt.Next = curr
				gt = curr
			} else {
				gtHead = curr // head of greater numbers
				gt = curr
			}
		}

		curr = curr.Next // next number
	}

	if ltHead == nil {
		return gtHead
	}

	lt.Next = gtHead // concatenate it
	if gt != nil {
		gt.Next = nil
	}
	return ltHead
}

// helper
func makeListNode(a []int) *ListNode {
	head := &ListNode{Val: a[0]}

	curr := head
	for i := 1; i < len(a); i++ {
		tmp := &ListNode{Val: a[i]}
		curr.Next = tmp
		curr = tmp
	}
	return head
}

func TestPartition(t *testing.T) {

	lst := partition(makeListNode([]int{1, 4, 3, 2, 5, 2}), 3)

	for lst != nil {
		fmt.Printf("%d ", lst.Val)
		lst = lst.Next
	}
}

// https://leetcode.com/problems/middle-of-the-linked-list

package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	// node6 := &ListNode{Val: 6}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	// node5.Next = node6

	fmt.Println(*middleNode(node1))
}

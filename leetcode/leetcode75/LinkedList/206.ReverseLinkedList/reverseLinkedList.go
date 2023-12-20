// https://leetcode.com/problems/reverse-linked-list/

package main

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

func reverseList(head *ListNode) *ListNode {
	result := &ListNode{}
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = result.Next
		result.Next = curr
		curr = next
	}
	return result.Next
}

//
// Author: Duc Tran
// Title: Two Sum
//

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy_head := &ListNode{}
	current, carry := dummy_head, 0

	for l1 != nil || l2 != nil {
		val := carry

		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry, val = val%10, val%10
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return dummy_head.Next
}

func main() {
	// list1 := &ListNode{
	// 	Val: 0,
	// 	Next: &ListNode{
	// 		Val:  8,
	// 		Next: nil,
	// 	},
	// }
	// list2 := &ListNode{
	// 	Val:  0,
	// 	Next: nil,
	// }
	// result := addTwoNumbers(list1, list2)
	// fmt.Println(result.Val)
}

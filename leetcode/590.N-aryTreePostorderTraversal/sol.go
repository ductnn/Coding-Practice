package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	for _, child := range root.Children {
		res = append(res, postorder(child)...)
	}
	res = append(res, root.Val)

	return res
}

func main() {
	root := &Node{Val: 1}
	child1 := &Node{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 4}
	root.Children = []*Node{child1, child2, child3}

	result := postorder(root)
	fmt.Println(result)
}

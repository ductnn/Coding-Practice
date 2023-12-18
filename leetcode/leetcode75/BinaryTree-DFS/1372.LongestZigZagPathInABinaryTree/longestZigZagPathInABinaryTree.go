package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	cnt := 0
	var dfs func(root *TreeNode, left, right int)
	dfs = func(root *TreeNode, left, right int) {
		if root == nil {
			return
		}
		cnt = max(cnt, max(left, right))
		dfs(root.Left, right+1, 0)
		dfs(root.Right, 0, left+1)
	}
	dfs(root, 0, 0)
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 6},
			},
		},
	}

	fmt.Println(longestZigZag(root))
}

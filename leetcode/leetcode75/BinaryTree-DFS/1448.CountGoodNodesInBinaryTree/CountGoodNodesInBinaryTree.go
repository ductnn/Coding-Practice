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

func goodNodes(root *TreeNode) int {
	cnt := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, mx int) {
		if root == nil {
			return
		}
		if mx <= root.Val {
			cnt++
			mx = root.Val
		}

		dfs(root.Left, mx)
		dfs(root.Right, mx)
	}

	dfs(root, -10001)
	return cnt
}

func goodNodes1(root *TreeNode) int {
	return rec(root.Left, root.Right, root.Val) + 1
}

func rec(rootLeft *TreeNode, rootRight *TreeNode, mx int) int {
	cnt := 0

	if rootLeft != nil && rootLeft.Val >= mx {
		cnt++
	}
	if rootLeft != nil {
		cnt += rec(rootLeft.Left, rootLeft.Right, max(mx, rootLeft.Val))
	}

	if rootRight != nil && rootRight.Val >= mx {
		cnt++
	}
	if rootRight != nil {
		cnt += rec(rootRight.Left, rootRight.Right, max(mx, rootRight.Val))
	}

	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Define the trees
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 5},
		},
	}

	// Check leaf similarity
	fmt.Println(goodNodes1(root))
}

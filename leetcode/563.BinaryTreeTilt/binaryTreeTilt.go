package main

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

func findTilt(root *TreeNode) int {
	result := 0
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left, result), dfs(root.Right, result)
	*result += abs(left - right)
	return root.Val + right + left
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {}

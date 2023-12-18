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

func pathSum(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum += root.Val
		result := cnt[sum-targetSum]
		cnt[sum]++
		result += dfs(root.Left, sum) + dfs(root.Right, sum)
		cnt[sum]--
		return result
	}

	return dfs(root, 0)
}

func main() {
	// Define the trees
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Right: &TreeNode{Val: 11},
		},
	}
	targetSum := 8

	fmt.Println(pathSum(root, targetSum))
}

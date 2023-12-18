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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := []int{}
	dfs(root1, &arr1)

	arr2 := []int{}
	dfs(root2, &arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
		return
	}
	dfs(root.Left, arr)
	dfs(root.Right, arr)
}

func main() {
	// Define the trees
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 9},
		},
	}

	// Check leaf similarity
	result := leafSimilar(root1, root2)
	fmt.Println(result)
}

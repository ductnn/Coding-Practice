// https://leetcode.com/problems/combination-sum/

package main

import (
	"fmt"
	"slices"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	temp := []int{}
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if s == 0 {
			result = append(result, slices.Clone(temp))
			return
		}
		if s < candidates[i] {
			return
		}
		for j := i; j < len(candidates); j++ {
			temp = append(temp, candidates[j])
			dfs(j, s-candidates[j])
			temp = temp[:len(temp)-1]
		}

	}
	dfs(0, target)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinationSum(candidates, target))
}

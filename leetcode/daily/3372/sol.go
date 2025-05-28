// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/description/

package main

import "fmt"

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n1 := 0
	n2 := 0

	for _, edge := range edges1 {
		if edge[0] > n1 {
			n1 = edge[0]
		}
		if edge[1] > n1 {
			n1 = edge[1]
		}
	}

	for _, edge := range edges2 {
		if edge[0] > n2 {
			n2 = edge[0]
		}
		if edge[1] > n2 {
			n2 = edge[1]
		}
	}

	graph1 := make([][]int, n1+1)
	graph2 := make([][]int, n2+1)

	for _, edge := range edges1 {
		graph1[edge[0]] = append(graph1[edge[0]], edge[1])
		graph1[edge[1]] = append(graph1[edge[1]], edge[0])
	}

	for _, edge := range edges2 {
		graph2[edge[0]] = append(graph2[edge[0]], edge[1])
		graph2[edge[1]] = append(graph2[edge[1]], edge[0])
	}

	maxTargetsTree2 := 0
	for i := 0; i <= n2; i++ {
		reachable := dfs(graph2, i, -1, k-1)
		if reachable > maxTargetsTree2 {
			maxTargetsTree2 = reachable
		}
	}

	result := make([]int, n1+1)
	for i := 0; i <= n1; i++ {
		reachableInTree1 := dfs(graph1, i, -1, k)
		result[i] = reachableInTree1 + maxTargetsTree2
	}

	return result
}

func dfs(graph [][]int, node int, parent int, depth int) int {
	if depth < 0 {
		return 0
	}
	count := 1
	for _, neighbor := range graph[node] {
		if neighbor != parent {
			count += dfs(graph, neighbor, node, depth-1)
		}
	}
	return count
}

func main() {
	edges1 := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	edges2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}
	k := 2

	fmt.Println(maxTargetNodes(edges1, edges2, k))
}

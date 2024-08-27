package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node int
	prob float64
}

type MaxHeap []Edge

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build graph
	graph := make([][]Edge, n)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		prob := succProb[i]
		graph[u] = append(graph[u], Edge{v, prob})
		graph[v] = append(graph[v], Edge{u, prob})
	}

	// Max-heap for Dijkstra's algorithm
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Edge{start, 1.0})

	// Array to store the max probability to reach each node
	maxProb := make([]float64, n)
	maxProb[start] = 1.0

	for maxHeap.Len() > 0 {
		curr := heap.Pop(maxHeap).(Edge)
		currNode, currProb := curr.node, curr.prob

		// If we reach the end node, return the probability
		if currNode == end {
			return currProb
		}

		// Update probabilities for all adjacent nodes
		for _, edge := range graph[currNode] {
			nextNode, nextProb := edge.node, edge.prob
			// Calculate the new probability
			newProb := currProb * nextProb
			if newProb > maxProb[nextNode] {
				maxProb[nextNode] = newProb
				heap.Push(maxHeap, Edge{nextNode, newProb})
			}
		}
	}

	// If the end node is not reachable, return 0
	return 0.0
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2

	result := maxProbability(n, edges, succProb, start, end)
	fmt.Println(result) // Output: 0.25
}

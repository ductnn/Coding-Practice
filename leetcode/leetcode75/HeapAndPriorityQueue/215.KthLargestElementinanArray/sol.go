// https://leetcode.com/problems/kth-largest-element-in-an-array

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// with sorting
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// without sorting
type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] <= h[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest1(nums []int, k int) int {
	tmp := Heap(nums)
	h := &tmp
	heap.Init(h)
	for len(*h) > k {
		heap.Pop(h)
	}
	return (*h)[0]
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	fmt.Println(findKthLargest1(nums, k))
}

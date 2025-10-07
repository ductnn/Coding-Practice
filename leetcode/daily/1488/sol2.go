// https://leetcode.com/problems/avoid-flood-in-the-city

package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	day  int
	lake int
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].day < h[j].day }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	full := make(map[int]bool)  // hồ nào đang đầy
	next := make(map[int][]int) // hồ -> danh sách ngày mưa tiếp theo

	// Tiền xử lý: lưu trước các ngày mưa của từng hồ
	for i, lake := range rains {
		if lake > 0 {
			next[lake] = append(next[lake], i)
		}
	}

	h := &MinHeap{}
	heap.Init(h)

	for i, lake := range rains {
		if lake > 0 {
			// Bỏ qua ngày hôm nay
			next[lake] = next[lake][1:]

			if full[lake] {
				// Mưa vào hồ đang đầy -> flood
				return []int{}
			}
			full[lake] = true
			ans[i] = -1

			// Nếu còn mưa hồ này trong tương lai -> đưa vào heap
			if len(next[lake]) > 0 {
				heap.Push(h, Pair{day: next[lake][0], lake: lake})
			}
		} else {
			// Ngày khô
			if h.Len() > 0 {
				p := heap.Pop(h).(Pair)
				ans[i] = p.lake
				full[p.lake] = false
			} else {
				ans[i] = 1 // dry hồ bất kỳ
			}
		}
	}

	return ans
}

func main() {
	rains := []int{1, 2, 0, 0, 2, 1}
	fmt.Println(avoidFlood(rains)) // expected: [-1, -1, 2, 1, -1, -1]
}

// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
package main

func queryResults(limit int, queries [][]int) (ans []int) {
	g := map[int]int{}
	cnt := map[int]int{}
	for _, q := range queries {
		x, y := q[0], q[1]
		cnt[y]++
		if v, ok := g[x]; ok {
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
		g[x] = y
		ans = append(ans, len(cnt))
	}
	return
}

func main() {
	queries := [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}
	limit := 4
	println(queryResults(limit, queries))
}

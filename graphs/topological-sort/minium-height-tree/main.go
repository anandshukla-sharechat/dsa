package main

import (
	"container/list"
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	inDegree := make([]int, n)
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	q := list.New()
	for _, val := range adj {
		for _, v := range val {
			inDegree[v]++
		}
	}

	for i := 0; i < n; i++ {
		if inDegree[i] == 1 {
			q.PushBack(i)
		}
	}
	ans := make([]int, 0)

	for q.Len() > 0 {
		sz := q.Len()
		if len(ans) > 0 {
			ans = ans[:0]
		}
		for sz > 0 {
			el := q.Remove(q.Front()).(int)
			ans = append(ans, el)
			for _, val := range adj[el] {
				inDegree[val]--
				if inDegree[val] == 1 {
					q.PushBack(val)
				}
			}
			sz--
		}
	}
	return ans
}

// https://leetcode.com/problems/minimum-height-trees/
func main() {
	n := 4
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	ans := findMinHeightTrees(n, edges)
	fmt.Println(ans)

}

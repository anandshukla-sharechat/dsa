package main

import (
	"container/list"
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	degree := make(map[int]int)
	adj := make(map[int][]int)

	for i := 0; i < len(graph); i++ {
		degree[i] = 0
		adj[i] = make([]int, 0)
	}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			adj[graph[i][j]] = append(adj[graph[i][j]], i)
		}
	}
	for _, arr := range adj {
		for _, val := range arr {
			degree[val]++
		}
	}
	q := list.New()
	for key, freq := range degree {
		if freq == 0 {
			q.PushBack(key)
		}
	}
	ans := make([]int, 0)
	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			el := q.Remove(q.Front()).(int)
			ans = append(ans, el)
			for _, val := range adj[el] {
				degree[val]--
				if degree[val] == 0 {
					q.PushBack(val)
				}
			}
			sz--
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	ans := eventualSafeNodes(graph)
	fmt.Println(ans)
}

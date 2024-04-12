package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	neighbour int
	weight    int
}

type MinHeap []Node

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Node))
}

func (m *MinHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func prepareAdj(n int, edges [][]int) map[int][]Node {
	mp := make(map[int][]Node)
	for i := 1; i <= n; i++ {
		mp[i] = make([]Node, 0)
	}
	for _, val := range edges {
		mp[val[0]] = append(mp[val[0]], Node{
			neighbour: val[1],
			weight:    val[2],
		})
		mp[val[1]] = append(mp[val[1]], Node{
			neighbour: val[0],
			weight:    val[2],
		})
	}
	return mp
}

func dfs(node int, cost map[int]int, visited []bool, ans *int, adj map[int][]Node, n int) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, val := range adj[node] {
		if cost[node] < cost[val.neighbour] {
			if val.neighbour == 1 {
				*ans = (*ans + 1) % (1e9 + 7)
				continue
			}
			dfs(val.neighbour, cost, visited, ans, adj, n)
		}
	}
	visited[node] = false
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := prepareAdj(n, edges)
	pq := make(MinHeap, 0)
	heap.Init(&pq)
	cost := make(map[int]int)
	vis := make(map[int]bool)
	for i := 1; i <= n; i++ {
		cost[i] = math.MaxInt
		vis[i] = false
	}
	heap.Push(&pq, Node{neighbour: n, weight: 0})
	cost[n] = 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Node)
		if vis[u.neighbour] {
			continue
		}
		for _, val := range adj[u.neighbour] {
			if cost[val.neighbour] > (cost[u.neighbour]+val.weight) && !vis[val.neighbour] {
				cost[val.neighbour] = cost[u.neighbour] + val.weight
				heap.Push(&pq, Node{
					neighbour: val.neighbour,
					weight:    cost[val.neighbour],
				})
			}
		}
		vis[u.neighbour] = true
	}
	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		visited[i] = false
	}

	ans := 0
	dfs(n, cost, visited, &ans, adj, n)
	return ans
}

/*
52 / 77 testcases passed
https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/description/
*/
func main() {
	n, edges := 5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}
	ans := countRestrictedPaths(n, edges)
	fmt.Println(ans)
}

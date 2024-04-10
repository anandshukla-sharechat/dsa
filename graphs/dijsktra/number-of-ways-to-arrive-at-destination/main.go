package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	intersections int
	time          int
}

type Minheap []Node

func (m Minheap) Len() int {
	return len(m)
}

func (m Minheap) Less(i, j int) bool {
	return m[i].time < m[j].time
}

func (m Minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *Minheap) Push(x any) {
	*m = append(*m, x.(Node))
}

func (m *Minheap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func prepareAdj(n int, roads [][]int) map[int][]Node {
	adj := make(map[int][]Node)
	for i := 0; i < n; i++ {
		adj[i] = make([]Node, 0)
	}
	for i := 0; i < len(roads); i++ {
		adj[roads[i][0]] = append(adj[roads[i][0]], Node{
			intersections: roads[i][1],
			time:          roads[i][2],
		})
		adj[roads[i][1]] = append(adj[roads[i][1]], Node{
			intersections: roads[i][0],
			time:          roads[i][2],
		})
	}
	return adj
}

func countPaths(n int, roads [][]int) int {
	mod := int(1e9 + 7)
	adj := prepareAdj(n, roads)
	pq := make(Minheap, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt
	}
	mp := make(map[int]int)
	heap.Init(&pq)
	heap.Push(&pq, Node{0, 0})
	cost[0] = 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Node)
		for _, val := range adj[u.intersections] {
			newCost := cost[u.intersections] + val.time
			if cost[val.intersections] >= newCost {
				cost[val.intersections] = newCost
				heap.Push(&pq, Node{
					intersections: val.intersections,
					time:          cost[val.intersections],
				})
			}
			if val.intersections == (n - 1) {
				mp[newCost] = (mp[newCost] + 1) % mod
			}
		}
	}
	return mp[cost[n-1]]
}

// 20 / 55 testcases passed
// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description/
func main() {
	n, roads := 7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}
	ans := countPaths(n, roads)
	fmt.Println(ans)
}

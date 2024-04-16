package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	target int
	weight int
}

type Minheap []Node

func (m Minheap) Len() int {
	return len(m)
}

func (m Minheap) Less(i, j int) bool {
	return m[i].weight < m[j].weight
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

func prepareAdj(n int, times [][]int) map[int][]Node {
	mp := make(map[int][]Node)
	for i := 1; i <= n; i++ {
		mp[i] = make([]Node, 0)
	}
	for _, val := range times {
		mp[val[0]] = append(mp[val[0]], Node{target: val[1], weight: val[2]})
	}
	return mp
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := prepareAdj(n, times)
	pq := make(Minheap, 0)
	heap.Init(&pq)
	cost := make(map[int]int)
	visited := make(map[int]bool)
	for i := 1; i <= n; i++ {
		cost[i] = math.MaxInt
		visited[i] = false
	}
	heap.Push(&pq, Node{target: k, weight: 0})
	cost[k] = 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Node)
		for _, val := range adj[u.target] {
			if !visited[val.target] && cost[val.target] > (cost[u.target]+val.weight) {
				cost[val.target] = cost[u.target] + val.weight
				heap.Push(&pq, Node{target: val.target, weight: cost[val.target]})
			}
		}
		visited[u.target] = true
	}
	ans := math.MinInt
	for i := 1; i <= n; i++ {
		ans = max(ans, cost[i])
	}
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}

// https://leetcode.com/problems/network-delay-time/description/
func main() {
	times, n, k := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2
	ans := networkDelayTime(times, n, k)
	fmt.Println(ans)
}

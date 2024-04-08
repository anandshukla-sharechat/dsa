package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	node int
	wt   float64
}

type MaxHeap []node

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].wt > m[j].wt
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(node))
}

func (m *MaxHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func buildAdjacencyList(n int, edges [][]int, succProb []float64) [][]node {
	adj := make([][]node, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]node, 0)
	}
	for i, val := range edges {
		adj[val[0]] = append(adj[val[0]], node{node: val[1], wt: succProb[i]})
		adj[val[1]] = append(adj[val[1]], node{node: val[0], wt: succProb[i]})
	}
	return adj
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	adj := buildAdjacencyList(n, edges, succProb)
	visited := make([]bool, n)
	cost := make([]float64, n)
	for i := 0; i < n; i++ {
		visited[i] = true
		cost[i] = 0
	}
	pq := make(MaxHeap, 0)
	heap.Init(&pq)
	heap.Push(&pq, node{node: start_node, wt: 0})
	cost[start_node] = 1
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(node)
		for _, val := range adj[u.node] {
			if cost[val.node] < (cost[u.node] * val.wt) {
				cost[val.node] = cost[u.node] * val.wt
				heap.Push(&pq, node{
					node: val.node,
					wt:   cost[val.node],
				})
			}
		}
	}
	return cost[end_node]
}

// https://leetcode.com/problems/path-with-maximum-probability/description/
func main() {
	n, edges, succProb, start_node, end_node := 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2
	fmt.Println(maxProbability(n, edges, succProb, start_node, end_node))
}

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	n     int
	price int
	k     int
}

type MinHeap []Node

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	if m[i].k == m[j].k {
		return m[i].price < m[j].price
	}
	return m[i].k < m[j].k
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

func prepareAdjacenyList(n int, flights [][]int) map[int][]Node {
	mp := make(map[int][]Node)
	for i := 0; i < n; i++ {
		mp[i] = make([]Node, 0)
	}
	for _, val := range flights {
		mp[val[0]] = append(mp[val[0]], Node{
			n:     val[1],
			price: val[2],
		})
	}
	return mp
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := prepareAdjacenyList(n, flights)
	pq := make(MinHeap, 0)
	heap.Init(&pq)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt
	}
	heap.Push(&pq, Node{
		n:     src,
		price: 0,
		k:     0,
	})
	cost[src] = 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Node)
		if u.k <= k {
			for _, val := range adj[u.n] {
				val.price += u.price
				val.k = u.k + 1
				if cost[val.n] > val.price {
					heap.Push(&pq, val)
					cost[val.n] = val.price
				}
			}
		}
	}
	if cost[dst] == math.MaxInt {
		return -1
	} else {
		return cost[dst]
	}
}

// https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
func main() {
	n, flights, src, dst, k := 4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1
	ans := findCheapestPrice(n, flights, src, dst, k)
	fmt.Println(ans)
}

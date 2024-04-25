package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Node   int
	Weight int
	Time   int
}

type Minheap []Node

func (m Minheap) Len() int {
	return len(m)
}

func (m Minheap) Less(i, j int) bool {
	if m[i].Weight == m[j].Weight {
		return m[i].Time < m[j].Time
	} else {
		return m[i].Weight < m[j].Weight
	}
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

func prepareAdj(edges [][]int, n int) map[int][]Node {
	mp := make(map[int][]Node)
	for i := 0; i < n; i++ {
		mp[i] = make([]Node, 0)
	}
	for _, val := range edges {
		mp[val[0]] = append(mp[val[0]], Node{
			Node: val[1],
			Time: val[2],
		})
		mp[val[1]] = append(mp[val[1]], Node{
			Node: val[0],
			Time: val[2],
		})
	}
	return mp
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	adj := prepareAdj(edges, n)
	pq := make(Minheap, 0)
	time := make([]int, n)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		time[i] = math.MaxInt
		cost[i] = math.MaxInt
	}
	heap.Init(&pq)
	cost[0] = passingFees[0]
	time[0] = 0
	heap.Push(&pq, Node{
		Node:   0,
		Weight: cost[0],
		Time:   time[0],
	})
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Node)
		for _, val := range adj[u.Node] {
			if u.Time+val.Time <= maxTime {
				if cost[val.Node] > (u.Weight + passingFees[val.Node]) {
					cost[val.Node] = u.Weight + passingFees[val.Node]
					time[val.Node] = time[u.Node] + val.Time
					heap.Push(&pq, Node{
						Node:   val.Node,
						Time:   time[val.Node],
						Weight: cost[val.Node],
					})
				} else if time[val.Node] > (u.Time + val.Time) {
					time[val.Node] = u.Time + val.Time
					heap.Push(&pq, Node{
						Node:   val.Node,
						Time:   time[val.Node],
						Weight: u.Weight + passingFees[val.Node],
					})
				}
			}
		}
	}
	if cost[n-1] == math.MaxInt {
		return -1
	}
	return cost[n-1]
}

// https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/description/
func main() {
	maxTime, edges, passingFees := 30, [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}, []int{5, 1, 2, 20, 20, 3}
	ans := minCost(maxTime, edges, passingFees)
	fmt.Println(ans)
}

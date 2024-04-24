package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Node     int
	Weight   int
	Discount int
}

type Minheap []Node

func (m Minheap) Len() int {
	return len(m)
}

func (m Minheap) Less(i, j int) bool {
	return m[i].Weight < m[j].Weight
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

func prepareAdj(n int, highways [][]int) map[int][]Node {
	mp := make(map[int][]Node)
	for i := 0; i < n; i++ {
		mp[i] = make([]Node, 0)
	}
	for _, el := range highways {
		mp[el[0]] = append(mp[el[0]], Node{
			Node:   el[1],
			Weight: el[2],
		})
		mp[el[1]] = append(mp[el[1]], Node{
			Node:   el[0],
			Weight: el[2],
		})
	}
	return mp
}

func minimumCost(n int, highways [][]int, discounts int) int {
	adj := prepareAdj(n, highways)
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, discounts+1)
		for j := 0; j <= discounts; j++ {
			cost[i][j] = math.MaxInt
		}
	}
	pq := make(Minheap, 0)
	heap.Init(&pq)
	heap.Push(&pq, Node{Node: 0, Weight: 0, Discount: discounts})
	cost[0][discounts] = 0
	for pq.Len() > 0 {
		el := heap.Pop(&pq).(Node)
		for _, val := range adj[el.Node] {
			// without discount
			if cost[val.Node][el.Discount] > (cost[el.Node][el.Discount] + val.Weight) {
				cost[val.Node][el.Discount] = cost[el.Node][el.Discount] + val.Weight
				heap.Push(&pq, Node{
					Node:     val.Node,
					Weight:   cost[val.Node][el.Discount],
					Discount: el.Discount,
				})
			}
			// with discount
			if el.Discount > 0 && cost[val.Node][el.Discount-1] > (cost[el.Node][el.Discount]+val.Weight/2) {
				cost[val.Node][el.Discount-1] = cost[el.Node][el.Discount] + val.Weight/2
				heap.Push(&pq, Node{
					Node:     val.Node,
					Weight:   cost[val.Node][el.Discount-1],
					Discount: el.Discount - 1,
				})
			}
		}
	}
	ans := math.MaxInt
	for i := 0; i <= discounts; i++ {
		ans = min(ans, cost[n-1][i])
	}
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}

// https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/description/
/*
A series of highways connect n cities numbered from 0 to n - 1. You are given a 2D integer array highways where highways[i] = [city1i, city2i, tolli] indicates that there is a highway that connects city1i and city2i, allowing a car to go from city1i to city2i and vice versa for a cost of tolli.

You are also given an integer discounts which represents the number of discounts you have. You can use a discount to travel across the ith highway for a cost of tolli / 2 (integer division). Each discount may only be used once, and you can only use at most one discount per highway.

Return the minimum total cost to go from city 0 to city n - 1, or -1 if it is not possible to go from city 0 to city n - 1.
*/
func main() {
	n, highways, discount := 5, [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2}}, 1
	ans := minimumCost(n, highways, discount)
	fmt.Println(ans)
}

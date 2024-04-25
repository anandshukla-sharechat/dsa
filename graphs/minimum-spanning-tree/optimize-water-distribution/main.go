package main

import (
	"fmt"
	"sort"
)

type Node struct {
	a int
	b int
	w int
}

func find(x int, mp map[int]int) int {
	if mp[x] == x {
		return x
	} else {
		res := find(mp[x], mp)
		mp[x] = res
		return res
	}
}

func union(x, y int, mp map[int]int, group *int) {
	a := find(x, mp)
	b := find(y, mp)
	if a != b {
		mp[a] = b
		(*group)--
	}
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	mp := make(map[int]int)
	for i := 0; i <= n; i++ {
		mp[i] = i
	}
	edges := make([]Node, 0)
	// prepares edges from well
	for i := 0; i < len(wells); i++ {
		edges = append(edges, Node{a: 0, b: i + 1, w: wells[i]})
		edges = append(edges, Node{a: i + 1, b: 0, w: wells[i]})
	}
	// prepare edges from pipes
	for _, val := range pipes {
		edges = append(edges, Node{a: val[0], b: val[1], w: val[2]})
		edges = append(edges, Node{a: val[1], b: val[0], w: val[2]})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	groups := n + 1
	cost := 0

	for _, val := range edges {
		if find(val.a, mp) != find(val.b, mp) {
			cost += val.w
			union(val.a, val.b, mp, &groups)
		}
		if groups == 1 {
			return cost
		}
	}
	return -1
}

/*
https://leetcode.com/problems/optimize-water-distribution-in-a-village/description/
There are n houses in a village. We want to supply water for all the houses by building wells and laying pipes.

For each house i, we can either build a well inside it directly with cost wells[i - 1] (note the -1 due to 0-indexing), or pipe in water from another well to it. The costs to lay pipes between houses are given by the array pipes where each pipes[j] = [house1j, house2j, costj] represents the cost to connect house1j and house2j together using a pipe. Connections are bidirectional, and there could be multiple valid connections between the same two houses with different costs.

Return the minimum total cost to supply water to all houses.
*/
func main() {
	n, wells, pipes := 3, []int{1, 2, 2}, [][]int{{1, 2, 1}, {2, 3, 1}}
	ans := minCostToSupplyWater(n, wells, pipes)
	fmt.Println(ans)
}

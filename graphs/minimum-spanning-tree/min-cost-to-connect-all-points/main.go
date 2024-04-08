package main

import (
	"fmt"
	"sort"
)

func dist(a, b int) int {
	c := a - b
	if c < 0 {
		c = 0 - c
	}
	return c
}

type Node struct {
	a [2]int
	b [2]int
	w int
}

func find(x [2]int, mp map[[2]int][2]int) [2]int {
	if mp[x] == x {
		return x
	} else {
		res := find(mp[x], mp)
		mp[x] = res
		return res
	}
}

func union(x [2]int, y [2]int, mp map[[2]int][2]int) {
	a := find(x, mp)
	b := find(y, mp)
	if a != b {
		mp[a] = b
	}
}

func minCostConnectPoints(points [][]int) int {
	graph := make([]Node, 0)
	mp := make(map[[2]int][2]int)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			graph = append(graph, Node{
				a: [2]int{points[i][0], points[i][1]},
				b: [2]int{points[j][0], points[j][1]},
				w: dist(points[i][0], points[j][0]) + dist(points[i][1], points[j][1]),
			})
		}
		mp[[2]int{points[i][0], points[i][1]}] = [2]int{points[i][0], points[i][1]}
	}
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].w < graph[j].w
	})
	ans := 0
	for _, val := range graph {
		if find(val.a, mp) != find(val.b, mp) {
			union(val.a, val.b, mp)
			ans += val.w
		}
	}
	return ans
}

// https://leetcode.com/problems/min-cost-to-connect-all-points/description/
func main() {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(minCostConnectPoints(points))
}

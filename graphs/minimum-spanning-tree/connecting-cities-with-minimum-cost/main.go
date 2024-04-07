package main

import (
	"fmt"
	"sort"
)

type Conn struct {
	cost int
	one  int
	two  int
}

func find(i int, mp map[int]int) int {
	if mp[i] == i {
		return i
	} else {
		res := find(mp[i], mp)
		mp[i] = res
		return res
	}
}

func union(i, j int, mp map[int]int) {
	a := find(i, mp)
	b := find(j, mp)
	if a != b {
		mp[a] = b
	}
}

func minimumCost(n int, connections [][]int) int {
	graph := make([]Conn, 0)
	for _, val := range connections {
		graph = append(graph, Conn{
			cost: val[2],
			one:  val[0],
			two:  val[1],
		})
	}
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].cost < graph[j].cost
	})
	ans := 0
	mp := make(map[int]int)
	for _, val := range graph {
		mp[val.one] = val.one
		mp[val.two] = val.two
	}
	for _, val := range graph {
		if find(val.one, mp) != find(val.two, mp) {
			ans += val.cost
			union(val.one, val.two, mp)
		}
	}
	check := make(map[int]int)
	for i := 1; i <= n; i++ {
		check[find(i, mp)]++
	}
	if len(check) != 1 {
		return -1
	}
	return ans
}

// https://leetcode.com/problems/connecting-cities-with-minimum-cost/description/
func main() {
	n, connections := 3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	fmt.Println(minimumCost(n, connections))
}

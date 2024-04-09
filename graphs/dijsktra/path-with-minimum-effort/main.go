package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Cell struct {
	cell [2]int
	hd   int
}

type MinHeap []Cell

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].hd < m[j].hd
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Cell))
}

func (m *MinHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func heightDiff(a, b, x, y int, heights [][]int) int {
	res := heights[a][b] - heights[x][y]
	if res < 0 {
		res = 0 - res
	}
	return res
}

func makeAdjacencyList(heights [][]int) map[[2]int][]Cell {
	n, m := len(heights), len(heights[0])
	dir := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	adj := make(map[[2]int][]Cell)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			adj[[2]int{i, j}] = make([]Cell, 0)
			for k := 0; k < 4; k++ {
				x := i + dir[k][0]
				y := j + dir[k][1]
				if x >= 0 && x < n && y >= 0 && y < m {
					adj[[2]int{i, j}] = append(adj[[2]int{i, j}], Cell{
						cell: [2]int{x, y},
						hd:   heightDiff(i, j, x, y, heights),
					})
				}
			}
		}
	}
	return adj
}

func minimumEffortPath(heights [][]int) int {
	adj := makeAdjacencyList(heights)
	n, m := len(heights), len(heights[0])
	pq := make(MinHeap, 0)
	heap.Init(&pq)
	cost := make(map[[2]int]int, n*m)
	visited := make(map[[2]int]bool, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cost[[2]int{i, j}] = math.MaxInt
		}
	}
	heap.Push(&pq, Cell{
		cell: [2]int{0, 0},
		hd:   0,
	})
	cost[[2]int{0, 0}] = 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(Cell)
		for _, val := range adj[u.cell] {
			if cost[val.cell] >= max(cost[u.cell], val.hd) && !visited[u.cell] {
				cost[val.cell] = max(cost[u.cell], val.hd)
				heap.Push(&pq, Cell{
					cell: val.cell,
					hd:   cost[val.cell],
				})
			}
		}
		visited[u.cell] = true
	}
	return cost[[2]int{n - 1, m - 1}]
}

// https://leetcode.com/problems/path-with-minimum-effort/description/
func main() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	ans := minimumEffortPath(heights)
	fmt.Println(ans)
}

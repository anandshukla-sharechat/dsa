package main

import (
	"container/list"
	"fmt"
)

type Coordinate struct {
	x int
	y int
}

func solve(m, n int, h [][]int, q *list.List, mark *[][]bool) {
	dir := []Coordinate{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			el := q.Remove(q.Front()).(Coordinate)
			(*mark)[el.x][el.y] = true
			for _, val := range dir {
				x := el.x + val.x
				y := el.y + val.y
				if x >= 0 && x < m && y >= 0 && y < n && h[x][y] >= h[el.x][el.y] && !(*mark)[x][y] {
					q.PushBack(Coordinate{x: x, y: y})
				}
			}

		}
	}
}

func pacificAtlantic(h [][]int) [][]int {
	m, n := len(h), len(h[0])
	q1, q2 := list.New(), list.New()
	m1, m2 := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		m1[i], m2[i] = make([]bool, n), make([]bool, n)
	}

	// Pacific Ocean
	for i := 0; i < m; i++ {
		q1.PushBack(Coordinate{x: i, y: 0})
	}
	for i := 1; i < n; i++ {
		q1.PushBack(Coordinate{x: 0, y: i})
	}

	// Atlantic Ocean
	for i := 0; i < m; i++ {
		q2.PushBack(Coordinate{x: i, y: n - 1})
	}
	for i := 0; i < n; i++ {
		q2.PushBack(Coordinate{x: m - 1, y: i})
	}
	solve(m, n, h, q1, &m1)
	solve(m, n, h, q2, &m2)

	ans := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if m1[i][j] && m2[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/pacific-atlantic-water-flow/
func main() {
	h := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	ans := pacificAtlantic(h)
	fmt.Println(ans)
}

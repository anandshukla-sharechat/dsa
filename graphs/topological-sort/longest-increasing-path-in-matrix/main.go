package main

import (
	"container/list"
	"fmt"
)

type coordinate struct {
	x int
	y int
}

func longestIncreasingPath(matrix [][]int) int {
	adj := make(map[coordinate][]coordinate)
	degree := make(map[coordinate]int)
	m, n := len(matrix), len(matrix[0])
	dir := []coordinate{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			degree[coordinate{
				x: i, y: j,
			}] = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 4; k++ {
				x := i + dir[k].x
				y := j + dir[k].y
				if x >= 0 && y >= 0 && x < m && y < n && matrix[i][j] < matrix[x][y] {
					if _, ok := adj[coordinate{x: x, y: y}]; ok {
						adj[coordinate{x: x, y: y}] = append(adj[coordinate{x: x, y: y}], coordinate{x: i, y: j})
					} else {
						adj[coordinate{x: x, y: y}] = []coordinate{{x: i, y: j}}
					}
				}
			}
		}
	}
	for _, coordinates := range adj {
		for _, vertex := range coordinates {
			degree[vertex]++
		}
	}

	ans := 0
	q := list.New()
	for xy, freq := range degree {
		if freq == 0 {
			q.PushBack(xy)
		}
	}

	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			el := q.Remove(q.Front()).(coordinate)
			for _, coordinates := range adj[el] {
				degree[coordinates]--
				if degree[coordinates] == 0 {
					q.PushBack(coordinates)
				}
			}
			sz--
		}
		ans++
	}

	return ans
}

func main() {
	matrix := [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}
	ans := longestIncreasingPath(matrix)
	fmt.Println(ans)
}

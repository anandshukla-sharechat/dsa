package main

import (
	"container/list"
	"fmt"
)

func constructSafeScoreGrid(grid [][]int) [][]int {
	n := len(grid)
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	safeGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		safeGrid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			safeGrid[i][j] = -1
		}
	}
	q := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q.PushBack([2]int{i, j})
				safeGrid[i][j] = 0
			}
		}
	}
	l := 1
	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			el := q.Remove(q.Front()).([2]int)
			for i := 0; i < len(dir); i++ {
				x := el[0] + dir[i][0]
				y := el[1] + dir[i][1]
				if x >= 0 && x < n && y >= 0 && y < n && safeGrid[x][y] == -1 {
					safeGrid[x][y] = l
					q.PushBack([2]int{x, y})
				}
			}
			sz--
		}
		l++
	}
	return safeGrid
}

func check(grid [][]int, mid int) bool {
	// apply bfs to check if mid satisfies lowest manhattan distance of safeness in path
	if mid > grid[0][0] {
		return false
	}
	n := len(grid)
	q := list.New()
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	vis := make(map[[2]int]int)
	q.PushBack([2]int{0, 0})
	vis[[2]int{0, 0}]++
	for q.Len() > 0 {
		el := q.Remove(q.Front()).([2]int)
		if el[0] == (n-1) && el[1] == (n-1) {
			return true
		}
		for i := 0; i < len(dir); i++ {
			x := el[0] + dir[i][0]
			y := el[1] + dir[i][1]
			if x >= 0 && x < n && y >= 0 && y < n && vis[[2]int{x, y}] == 0 && grid[x][y] >= mid {
				vis[[2]int{x, y}]++
				q.PushBack([2]int{x, y})
			}
		}
	}
	return false
}

func maximumSafenessFactor(grid [][]int) int {
	safeGrid := constructSafeScoreGrid(grid)

	l, u, ans := 0, int(1e8), -1
	for l <= u {
		mid := l + (u-l)/2
		if check(safeGrid, mid) {
			ans = mid
			l = mid + 1
		} else {
			u = mid - 1
		}
	}

	return ans
}

// https://leetcode.com/problems/find-the-safest-path-in-a-grid/description/
func main() {
	grid := [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}
	ans := maximumSafenessFactor(grid)
	fmt.Println(ans)
}

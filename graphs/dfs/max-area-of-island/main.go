package main

import "fmt"

func dfs(grid [][]int, visited *[][]bool, i, j, m, n int, area *int) {
	if i < 0 || i >= m {
		return
	}
	if j < 0 || j >= n {
		return
	}
	if grid[i][j] == 0 || (*visited)[i][j] {
		return
	}
	*area++
	(*visited)[i][j] = true
	dfs(grid, visited, i+1, j, m, n, area)
	dfs(grid, visited, i-1, j, m, n, area)
	dfs(grid, visited, i, j+1, m, n, area)
	dfs(grid, visited, i, j-1, m, n, area)
}
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	mark := make([][]bool, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area := 0
			if grid[i][j] == 1 && !mark[i][j] {
				dfs(grid, &mark, i, j, m, n, &area)
				ans = max(ans, area)
			}
		}
	}
	return ans
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	ans := maxAreaOfIsland(grid)
	fmt.Println(ans)
}

package main

import "fmt"

func dfs(grid [][]byte, visited *[][]bool, i, j, m, n int) {
	if i < 0 || i >= m {
		return
	}
	if j < 0 || j >= n {
		return
	}
	if grid[i][j] == '0' || (*visited)[i][j] {
		return
	}
	(*visited)[i][j] = true
	dfs(grid, visited, i+1, j, m, n)
	dfs(grid, visited, i-1, j, m, n)
	dfs(grid, visited, i, j+1, m, n)
	dfs(grid, visited, i, j-1, m, n)
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	mark := make([][]bool, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !mark[i][j] {
				dfs(grid, &mark, i, j, m, n)
				ans++
			}
		}
	}
	return ans
}

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	ans := numIslands(grid)
	fmt.Println(ans)
}

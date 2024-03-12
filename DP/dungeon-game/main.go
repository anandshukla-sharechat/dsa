package main

import "fmt"

/*
type Store struct {
	Value int
	Min   int
}

// 32 / 45 testcases passed

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	dp := make([][]Store, m)
	for i := range dp {
		dp[i] = make([]Store, n)
	}

	// top left cell
	dp[0][0].Value = dungeon[0][0]
	if dp[0][0].Value <= 0 {
		dp[0][0].Min = 1 - dp[0][0].Value
	} else {
		dp[0][0].Min = 0
	}

	// first row
	for i := 1; i < n; i++ {
		dp[0][i].Value = dp[0][i-1].Value + dungeon[0][i]
		if dp[0][i].Value <= 0 {
			dp[0][i].Min = max(1-dp[0][i].Value, dp[0][i-1].Min)
		} else {
			dp[0][i].Min = dp[0][i-1].Min
		}
	}

	// first column
	for i := 1; i < m; i++ {
		dp[i][0].Value = dp[i-1][0].Value + dungeon[i][0]
		if dp[i][0].Value <= 0 {
			dp[i][0].Min = max(1-dp[i][0].Value, dp[i-1][0].Min)
		} else {
			dp[i][0].Min = dp[i-1][0].Min
		}
	}

	// rest of processing
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fetchPreviousStore := func(a, b int) Store {
				if dp[a-1][b].Min < dp[a][b-1].Min {
					return dp[a-1][b]
				} else {
					return dp[a][b-1]
				}
			}
			previousStore := fetchPreviousStore(i, j)
			dp[i][j].Value = previousStore.Value + dungeon[i][j]
			if dp[i][j].Value <= 0 {
				dp[i][j].Min = max(1-dp[i][j].Value, previousStore.Min)
			} else {
				dp[i][j].Min = previousStore.Min
			}
		}
	}

	return max(dp[m-1][n-1].Min, 1)
}


*/

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]int, m)
	for i := range dp { // creating array
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1]) // 1 case

	for i := m - 2; i >= 0; i-- { //2 case
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(1, dp[m-1][i+1]-dungeon[m-1][i])
	}

	for i := m - 2; i >= 0; i-- { // 3 case
		for j := n - 2; j >= 0; j-- {
			tmp := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(1, tmp-dungeon[i][j])

		}
	}

	return dp[0][0]
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(dungeon))
}

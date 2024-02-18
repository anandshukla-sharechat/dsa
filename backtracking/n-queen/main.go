package main

import "fmt"

func isSafe(row, col, n int, grid [][]rune) bool {
	// left to right
	for i := 0; i < col; i++ {
		if grid[row][i] == 'Q' {
			return false
		}
	}
	// top left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if grid[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// bottom left diagonal
	for i, j := row+1, col-1; i < n && j >= 0; {
		if grid[i][j] == 'Q' {
			return false
		}
		i++
		j--
	}
	return true
}

func convertRuneTwoArrayToStringArray(grid [][]rune, n int) []string {
	temp := make([]string, n)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str = str + string(grid[i][j])
		}
		temp[i] = str
	}
	return temp
}

func nQueen(res *[][]string, grid [][]rune, col, n int) {
	if col == n {
		*res = append(*res, convertRuneTwoArrayToStringArray(grid, n))
		return
	}
	// iterate for rows
	for i := 0; i < n; i++ {
		if isSafe(i, col, n, grid) {
			grid[i][col] = 'Q'
			nQueen(res, grid, col+1, n)
			grid[i][col] = '.'
		}
	}
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	// prepare board
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		s := make([]rune, n)
		for j := 0; j < n; j++ {
			s[j] = '.'
		}
		board[i] = s
	}
	nQueen(&res, board, 0, n)
	return res
}

func printBoard(board []string) {
	for _, val := range board {
		fmt.Println(val)
	}
	fmt.Println("")
}

func main() {
	var n int
	fmt.Println("Enter the value of n ")
	fmt.Scanf("%d", &n)
	res := solveNQueens(n)
	for i := 0; i < len(res); i++ {
		printBoard(res[i])
	}
}

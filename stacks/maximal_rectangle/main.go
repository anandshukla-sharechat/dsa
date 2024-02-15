package main

import (
	"container/list"
	"fmt"
)

func findMaxRectangle(a []int) int {

	st := list.New()
	n := len(a)
	left, right := make([]int, n), make([]int, n)

	// left boundaries

	for i := 0; i < n; i = i + 1 {

		for st.Len() > 0 && a[st.Back().Value.(int)] >= a[i] {
			st.Remove(st.Back())
		}

		if st.Len() == 0 {
			left[i] = 0
		} else {
			left[i] = st.Back().Value.(int) + 1
		}
		st.PushBack(i)
	}

	// empty the stack
	for st.Len() > 0 {
		st.Remove(st.Back())
	}

	// right boundaries

	for i := n - 1; i >= 0; i = i - 1 {

		for st.Len() > 0 && a[st.Back().Value.(int)] >= a[i] {
			st.Remove(st.Back())
		}

		if st.Len() == 0 {
			right[i] = n - 1
		} else {
			right[i] = st.Back().Value.(int) - 1
		}
		st.PushBack(i)
	}

	// find max area
	ar := 0

	for i := 0; i < n; i = i + 1 {
		ar = max(ar, a[i]*(right[i]-left[i]+1))
	}

	return ar

}

func maximalRectangle(matrix [][]byte) int {

	m := len(matrix)
	n := len(matrix[0])

	a := make([][]int, m)
	for i := 0; i < m; i = i + 1 {
		a[i] = make([]int, n)
	}

	for i := 0; i < m; i = i + 1 {
		for j := 0; j < n; j = j + 1 {
			a[i][j] = int(matrix[i][j] - '0')
		}
	}

	dp := make([][]int, m)
	for i := 0; i < m; i = i + 1 {
		dp[i] = make([]int, n)
		copy(dp[i], a[i])
	}

	area := findMaxRectangle(dp[0])
	for i := 1; i < m; i = i + 1 {
		for j := 0; j < n; j = j + 1 {
			if dp[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j]
			}
		}
		area = max(area, findMaxRectangle(dp[i]))
	}
	return area

}

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}

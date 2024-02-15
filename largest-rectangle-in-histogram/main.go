package main

import (
	"container/list"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// https://leetcode.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	st := list.New()
	n := len(heights)
	leftArr, rightArr := make([]int, n), make([]int, n)

	// left small boundaries
	for i := 0; i < n; i = i + 1 {
		for st.Len() != 0 && heights[st.Back().Value.(int)] >= heights[i] {
			st.Remove(st.Back())
		}

		if st.Len() == 0 {
			leftArr[i] = 0
		} else {
			leftArr[i] = st.Back().Value.(int) + 1
		}
		st.PushBack(i)
	}

	// empty the stack
	for st.Len() > 0 {
		st.Remove(st.Back())
	}

	// right small boundaries
	for i := n - 1; i >= 0; i = i - 1 {
		for st.Len() != 0 && heights[st.Back().Value.(int)] >= heights[i] {
			st.Remove(st.Back())
		}
		if st.Len() == 0 {
			rightArr[i] = n - 1
		} else {
			rightArr[i] = st.Back().Value.(int) - 1
		}
		st.PushBack(i)
	}

	// find largest area
	area := 0
	for i := 0; i < n; i = i + 1 {
		area = max(area, heights[i]*(rightArr[i]-leftArr[i]+1))
	}
	return area
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

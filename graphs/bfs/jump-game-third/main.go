package main

import (
	"container/list"
	"fmt"
)

func canReach(arr []int, start int) bool {
	q := list.New()
	q.PushBack(start)
	mp, n := make(map[int]bool), len(arr)
	dir := []int{1, -1}
	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			frontEl := q.Remove(q.Front()).(int)
			mp[frontEl] = true
			if arr[frontEl] == 0 {
				return true
			}
			for _, v := range dir {
				pos := frontEl + v*arr[frontEl]
				if pos >= 0 && pos < n && !mp[pos] {
					q.PushBack(pos)
				}
			}
			sz--
		}
	}
	return false
}

// https://leetcode.com/problems/jump-game-iii/description/
func main() {
	arr, start := []int{4, 2, 3, 0, 3, 1, 2}, 5
	ans := canReach(arr, start)
	fmt.Println(ans)
}

package main

import (
	"container/list"
	"fmt"
)

func minJumps(arr []int) int {
	q := list.New()
	mp := make(map[int]bool)
	mp2 := make(map[int][]int)
	ans, n := 0, len(arr)
	q.PushBack(0)
	for idx, v := range arr {
		if _, ok := mp2[v]; !ok {
			mp2[v] = []int{idx}
		} else {
			mp2[v] = append(mp2[v], idx)
		}
	}
	for q.Len() > 0 {
		sz := q.Len()
		ans++
		for ; sz > 0; sz-- {
			frontEl := q.Remove(q.Front()).(int)
			if frontEl == (n - 1) {
				return ans - 1
			}
			if (frontEl-1) >= 0 && !mp[frontEl-1] {
				if (frontEl - 1) == (n - 1) {
					return ans
				}
				q.PushBack(frontEl - 1)
			}
			if (frontEl+1) < n && !mp[frontEl+1] {
				if (frontEl + 1) == (n - 1) {
					return ans
				}
				q.PushBack(frontEl + 1)
			}
			if !mp[frontEl] {
				for _, idx := range mp2[arr[frontEl]] {
					if !mp[idx] {
						q.PushBack(idx)
						mp[idx] = true
						if idx == (n - 1) {
							return ans
						}
					}
				}
			}
			mp[frontEl] = true
		}
	}
	return 0
}

// HARD : https://leetcode.com/problems/jump-game-iv/description/
func main() {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	ans := minJumps(arr)
	fmt.Println(ans)
}

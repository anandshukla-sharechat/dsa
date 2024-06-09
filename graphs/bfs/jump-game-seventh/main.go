package main

import (
	"container/list"
	"fmt"
)

func canReach(s string, minJump int, maxJump int) bool {
	q := list.New()
	mp, n, mx := make(map[int]bool), len(s), 0
	q.PushBack(0)

	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			frontEl := q.Remove(q.Front()).(int)
			mp[frontEl] = true
			if frontEl == (n-1) && (s[frontEl]-'0') == 0 {
				return true
			}
			for i := max(frontEl+minJump, mx+1); i <= min(frontEl+maxJump, n-1); i++ {
				if (s[i]-'0') == 0 && !mp[i] {
					q.PushBack(i)
					mp[i] = true
					if i == (n - 1) {
						return true
					}
				}
			}
			mx = max(mx, frontEl+maxJump)
		}
	}
	return false
}

/*
Input: s = "01101110", minJump = 2, maxJump = 3
Output: false
*/
// https://leetcode.com/problems/jump-game-vii/description/
func main() {
	s, minJump, maxJump := "011010", 2, 3
	ans := canReach(s, minJump, maxJump)
	fmt.Println(ans)
}

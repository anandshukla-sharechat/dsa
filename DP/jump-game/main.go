package main

import "fmt"

func canJump(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev == 0 {
			return false
		}
		prev = max(prev-1, nums[i])
	}
	return true
}

// https://leetcode.com/problems/jump-game/description/
func main() {
	nums := []int{2, 3, 1, 1, 4}
	ans := canJump(nums)
	fmt.Println(ans)
}

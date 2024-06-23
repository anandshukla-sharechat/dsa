package main

import "fmt"

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
		for j := 0; j < i; j++ {
			if abs(nums[i]-nums[j]) <= target && dp[j] >= 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func abs(a int) int {
	if a < 0 {
		a = 0 - a
		return a
	}
	return a
}

// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/
func main() {
	nums, target := []int{1, 3, 6, 4, 1, 2}, 2
	ans := maximumJumps(nums, target)
	fmt.Println(ans)
}

package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	} else if n == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	dp[n-1], dp[n-2], dp[n-3] = nums[n-1], nums[n-2], max(nums[n-3]+nums[n-1], nums[n-2])
	for i := n - 4; i >= 0; i-- {
		dp[i] = nums[i] + max(dp[i+2], dp[i+3])
		if i == 0 {
			dp[0] = max(dp[1], dp[0])
		}
	}
	return dp[0]
}

// https://leetcode.com/problems/house-robber/description/
func main() {
	nums := []int{1, 2, 3, 1}
	ans := rob(nums)
	fmt.Println(ans)
}

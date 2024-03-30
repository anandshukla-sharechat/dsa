package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	if amount == 0 {
		return 0
	}
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	// pre-processing for 1st iterations
	for _, coin := range coins {
		if coin < len(dp) {
			dp[coin] = 1
		}
	}
	for i := 1; i <= amount; i++ {

		for _, coin := range coins {
			if (i-coin) >= 0 && dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] != math.MaxInt {
		return dp[amount]
	} else {
		return -1
	}
}

/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0
*/
// https://leetcode.com/problems/coin-change/description/
func main() {
	coins := []int{1, 2, 5}
	amount := 11
	ans := coinChange(coins, amount)
	fmt.Println(ans)
}

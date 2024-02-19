package main

import (
	"fmt"
	"sort"
)

func solve(res *[][]int, sol []int, i int, nums []int, target int, su int) {
	if len(sol) >= 4 {
		if su == target {
			ans := make([]int, len(sol))
			copy(ans, sol)
			*res = append(*res, ans)
		}
		return
	}
	if i >= len(nums) {
		return
	}
	su += nums[i]
	sol = append(sol, nums[i])
	for j := i + 1; j < len(nums); j++ {
		su += nums[j]
		sol = append(sol, nums[j])
		solve(res, sol, j+1, nums, target, su)
		su -= nums[j]
		sol = sol[:len(sol)-1]
	}
	su -= nums[i]
	sol = sol[:len(sol)-1]
	solve(res, sol, i+1, nums, target, su)
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	sol := make([]int, 0)
	su := 0
	solve(&res, sol, 0, nums, target, su)

	tmp_res := make([][]int, 0)
	mp := make(map[[4]int]bool)
	for _, val := range res {
		temp_arr := [4]int{val[0], val[1], val[2], val[3]}
		temp_arr_2 := make([]int, 0)
		if !mp[temp_arr] {
			mp[temp_arr] = true
			for _, val := range temp_arr {
				temp_arr_2 = append(temp_arr_2, val)
			}
			tmp_res = append(tmp_res, temp_arr_2)
		}
	}
	return tmp_res
}

// https://leetcode.com/problems/4sum/submissions/1180256891/
// This solution passed 288 / 294 testcases, for rest test cases got TLE
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	res := fourSum(nums, target)
	for _, val := range res {
		fmt.Println(val)
	}
}

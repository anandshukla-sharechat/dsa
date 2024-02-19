package main

import "fmt"

func solve(res *[][]int, sol []int, i int, nums []int, target int, su int) {
	if su >= target {
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
	solve(res, sol, i, nums, target, su)
	su -= nums[i]
	sol = sol[:len(sol)-1]
	solve(res, sol, i+1, nums, target, su)
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	sol := make([]int, 0)
	su := 0
	solve(&res, sol, 0, candidates, target, su)
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	for _, val := range res {
		fmt.Println(val)
	}
}

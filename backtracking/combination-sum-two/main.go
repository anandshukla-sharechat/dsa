package main

import (
	"fmt"
	"sort"
)

func solve(candidates []int, target int, res *[][]int, sol []int, su int) {
	if su >= target {
		if su == target {
			temp := make([]int, len(sol))
			copy(temp, sol)
			*res = append(*res, temp)
		}
		return
	}
	if len(candidates) == 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		su += candidates[i]
		sol = append(sol, candidates[i])
		solve(candidates[i+1:], target, res, sol, su)
		su -= candidates[i]
		sol = sol[:len(sol)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sol := make([]int, 0)
	sort.Ints(candidates)
	solve(candidates, target, &res, sol, 0)
	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	for _, val := range res {
		fmt.Println(val)
	}
}

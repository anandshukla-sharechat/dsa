package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func solve(root *util.TreeNode, targetSum int, count *int, mp map[int]int, sum int) {
	if root == nil {
		return
	}
	sum += root.Val
	if sum == targetSum {
		*count = *count + 1
	}
	*count = *count + mp[sum-targetSum]
	mp[sum] += 1

	solve(root.Left, targetSum, count, mp, sum)
	solve(root.Right, targetSum, count, mp, sum)

	mp[sum] -= 1
}

func pathSum(root *util.TreeNode, targetSum int) int {
	count := 0
	mp := make(map[int]int)
	solve(root, targetSum, &count, mp, 0)
	return count
}

// https://leetcode.com/problems/path-sum-iii/description/
func main() {
	nums := []util.IntegerStruct{{10, true}, {5, true}, {-3, true}, {3, true}, {2, true}, {-1, false}, {11, true}, {3, true}, {-2, true}, {1, true}}
	root := util.LevelOrder(nums)
	output := pathSum(root, 8)
	fmt.Println(output)
}

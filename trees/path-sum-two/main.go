package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func solve(root *util.TreeNode, targetSum int, res *[][]int, sol []int, su int) {
	if root == nil {
		return
	}
	su += root.Val
	sol = append(sol, root.Val)
	if root.Left == nil && root.Right == nil && su == targetSum {
		temp := make([]int, len(sol))
		copy(temp, sol)
		*res = append(*res, temp)
	}
	solve(root.Left, targetSum, res, sol, su)
	solve(root.Right, targetSum, res, sol, su)
}

func pathSum(root *util.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	sol := make([]int, 0)
	solve(root, targetSum, &res, sol, 0)
	return res
}

func main() {
	nums := []util.IntegerStruct{{5, true}, {4, true}, {8, true}, {11, true}, {-1, false}, {13, true}, {4, true}, {7, true}, {2, true}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {5, true}, {1, true}}
	root := util.LevelOrder(nums)
	targetSum := 22
	res := pathSum(root, targetSum)
	for _, sol := range res {
		fmt.Println(sol)
	}
}

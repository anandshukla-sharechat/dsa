package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func hasPathSum(root *util.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	nums := []util.IntegerStruct{{5, true}, {4, true}, {8, true}, {11, true}, {-1, false}, {13, true}, {4, true}, {7, true}, {2, true}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {1, true}}
	root := util.LevelOrder(nums)
	targetSum := 22
	fmt.Println(hasPathSum(root, targetSum))
}

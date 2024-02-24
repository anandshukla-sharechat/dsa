package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
	"math"
)

func solve(root *util.TreeNode, minN int, maxN int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minN || root.Val >= maxN {
		return false
	}
	return solve(root.Left, minN, root.Val) && solve(root.Right, root.Val, maxN)
}

func isValidBST(root *util.TreeNode) bool {
	minN, maxN := math.MinInt, math.MaxInt
	return solve(root, minN, maxN)
}

func main() {
	treeArr := []util.IntegerStruct{{2, true}, {1, true}, {3, true}}
	root := util.LevelOrder(treeArr)
	fmt.Println(isValidBST(root))
}

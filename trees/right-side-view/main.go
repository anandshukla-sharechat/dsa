package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func solve(root *util.TreeNode, res *[]int, level *int, present_level int) {
	if root == nil {
		return
	}
	if (present_level + 1) == *level {
		*res = append(*res, root.Val)
		*level = *level + 1
	}
	solve(root.Right, res, level, present_level+1)
	solve(root.Left, res, level, present_level+1)
}

func rightSideView(root *util.TreeNode) []int {
	res := make([]int, 0)
	level := 1
	solve(root, &res, &level, 0)
	return res
}

func main() {
	nums := []util.IntegerStruct{{1, true}, {2, true}, {3, true}, {-1, false}, {5, true}, {-1, false}, {4, true}}
	root := util.LevelOrder(nums)
	ans := rightSideView(root)
	fmt.Println(ans)
}

package main

import (
	"container/list"
	"fmt"
	"github.com/anandshukla-sharechat/util"
	"slices"
)

func zigzagLevelOrder(root *util.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	sz := queue.Len()
	level := 1

	for sz > 0 {
		elements := make([]int, 0)
		for sz > 0 {
			el := queue.Remove(queue.Front()).(*util.TreeNode)
			elements = append(elements, el.Val)
			// for left child
			if el.Left != nil {
				queue.PushBack(el.Left)
			}
			if el.Right != nil {
				queue.PushBack(el.Right)
			}
			sz--
		}
		if level%2 == 0 {
			slices.Reverse(elements)
		}
		sz = queue.Len()
		ans = append(ans, elements)
		level++
	}
	return ans
}

func main() {
	root := util.LevelOrder([]util.IntegerStruct{{3, true}, {9, true}, {20, true}, {-1, false}, {-1, false}, {15, true}, {7, true}})
	ans := zigzagLevelOrder(root)
	for _, val := range ans {
		fmt.Println(val)
	}
}

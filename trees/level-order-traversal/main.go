package main

import (
	"container/list"
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func levelOrder(root *util.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	sz := queue.Len()

	for sz > 0 {
		levelEl := make([]int, 0)
		for sz > 0 {
			el := queue.Remove(queue.Front()).(*util.TreeNode)
			sz--
			levelEl = append(levelEl, el.Val)
			// left insertion
			if el.Left != nil {
				queue.PushBack(el.Left)
			}
			// right insertion
			if el.Right != nil {
				queue.PushBack(el.Right)
			}
		}
		sz = queue.Len()
		ans = append(ans, levelEl)
	}
	return ans
}

func main() {
	root := util.LevelOrder([]util.IntegerStruct{{3, true}, {9, true}, {20, true}, {-1, false}, {-1, false}, {15, true}, {7, true}})
	ans := levelOrder(root)
	for _, val := range ans {
		fmt.Println(val)
	}
}

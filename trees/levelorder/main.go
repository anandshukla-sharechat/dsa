package main

import (
	"container/list"
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

type IntegerStruct struct {
	Val     int
	NotNull bool
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	Node  *TreeNode
	Index int
}

func levelOrder(nums []IntegerStruct) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &TreeNode{
		Val:   nums[0].Val,
		Left:  nil,
		Right: nil,
	}
	queue := list.New()
	queue.PushBack(Element{
		Node:  head,
		Index: 0,
	})

	for queue.Len() > 0 {
		frontEl := queue.Front().Value.(Element)
		queue.Remove(queue.Front())

		// left child
		if (2*frontEl.Index+1) < n && nums[2*frontEl.Index+1].NotNull {
			leftChild := Element{
				Index: 2*frontEl.Index + 1,
				Node: &TreeNode{
					Val:   nums[2*frontEl.Index+1].Val,
					Left:  nil,
					Right: nil,
				},
			}
			frontEl.Node.Left = leftChild.Node
			queue.PushBack(leftChild)
		}

		// right child
		if (2*frontEl.Index+2) < n && nums[2*frontEl.Index+2].NotNull {
			rightChild := Element{
				Index: 2*frontEl.Index + 2,
				Node: &TreeNode{
					Val:   nums[2*frontEl.Index+2].Val,
					Left:  nil,
					Right: nil,
				},
			}
			frontEl.Node.Right = rightChild.Node
			queue.PushBack(rightChild)
		}
	}
	return head
}

func main() {
	nums := []util.IntegerStruct{{1, true}, {2, true}, {3, true}, {4, true}, {5, true}, {6, true}, {7, true}, {-1, false}, {8, true}, {9, true}}
	head := util.LevelOrder(nums)
	preorder := make([]int, 0)
	PreOrder(head, &preorder)
	fmt.Println(preorder)
	inorder := make([]int, 0)
	Inorder(head, &inorder)
	fmt.Println(inorder)
	postorder := make([]int, 0)
	PostOrder(head, &postorder)
	fmt.Println(postorder)
}

func PreOrder(root *util.TreeNode, ans *[]int) {
	if root != nil {
		*ans = append(*ans, root.Val)
	}
	if root.Left != nil {
		PreOrder(root.Left, ans)
	}
	if root.Right != nil {
		PreOrder(root.Right, ans)
	}
}

func Inorder(root *util.TreeNode, ans *[]int) {
	if root.Left != nil {
		Inorder(root.Left, ans)
	}
	if root != nil {
		*ans = append(*ans, root.Val)
	}
	if root.Right != nil {
		Inorder(root.Right, ans)
	}
}

func PostOrder(root *util.TreeNode, ans *[]int) {
	if root.Left != nil {
		PostOrder(root.Left, ans)
	}
	if root.Right != nil {
		PostOrder(root.Right, ans)
	}
	if root != nil {
		*ans = append(*ans, root.Val)
	}
}

package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
)

func computePath(root, x *util.TreeNode, path []*util.TreeNode, foundPath *[]*util.TreeNode) {
	if root == nil {
		return
	}
	path = append(path, root)
	if root == x {
		for _, el := range path {
			*foundPath = append(*foundPath, el)
		}
		return
	}
	computePath(root.Left, x, path, foundPath)
	computePath(root.Right, x, path, foundPath)
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
func lowestCommonAncestor(root, p, q *util.TreeNode) *util.TreeNode {

	pathOne, pathTwo := make([]*util.TreeNode, 0), make([]*util.TreeNode, 0)
	computePath(root, p, make([]*util.TreeNode, 0), &pathOne)
	computePath(root, q, make([]*util.TreeNode, 0), &pathTwo)
	var ans *util.TreeNode = nil
	for i := 0; i < len(pathOne) && i < len(pathTwo); i++ {
		if pathOne[i] == pathTwo[i] {
			ans = pathOne[i]
		} else {
			break
		}
	}

	return ans
}

func main() {
	root := util.LevelOrder([]util.IntegerStruct{{3, true}, {5, true}, {1, true}, {6, true}, {2, true}, {0, true}, {8, true}, {-1, false}, {-1, false}, {7, true}, {4, true}})
	ansNode := lowestCommonAncestor(root, root.Left, root.Right)
	fmt.Println(ansNode.Val)
}

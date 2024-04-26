package main

import (
	"fmt"
	"github.com/anandshukla-sharechat/util"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	serializeList         []string
	deserializedListIndex int
}

func Constructor() Codec {
	return Codec{serializeList: make([]string, 0), deserializedListIndex: 0}
}

func (this *Codec) preOrder(root *util.TreeNode) {
	if root == nil {
		this.serializeList = append(this.serializeList, "nil")
		return
	}
	this.serializeList = append(this.serializeList, fmt.Sprintf("%d", root.Val))
	this.preOrder(root.Left)
	this.preOrder(root.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *util.TreeNode) string {
	this.preOrder(root)
	serializedString := strings.Join(this.serializeList, ",")
	return serializedString
}

func (this *Codec) constructTreeFromPreOrderSerializedData(str []string) *util.TreeNode {
	if str[this.deserializedListIndex] == "nil" {
		return nil
	}
	val, err := strconv.Atoi(str[this.deserializedListIndex])
	if err != nil {
		return nil
	}
	root := &util.TreeNode{Val: val}
	this.deserializedListIndex++
	root.Left = this.constructTreeFromPreOrderSerializedData(str)
	this.deserializedListIndex++
	root.Right = this.constructTreeFromPreOrderSerializedData(str)
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *util.TreeNode {
	str := strings.Split(data, ",")
	root := this.constructTreeFromPreOrderSerializedData(str)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

func main() {
	nums := []util.IntegerStruct{{2, true}, {1, true}, {3, true}}
	root := util.LevelOrder(nums)
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(root)
	ans := deser.deserialize(tree)
	fmt.Println(ans)
}

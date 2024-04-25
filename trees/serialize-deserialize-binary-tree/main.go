package main

import (
	"container/list"
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
	deserializedList []*util.TreeNode
	serializedList   *list.List
}

func Constructor() Codec {
	return Codec{deserializedList: make([]*util.TreeNode, 0), serializedList: list.New()}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *util.TreeNode) string {
	ans := ""
	if root == nil {
		return ans
	}
	this.serializedList.PushBack(root)
	length := 0
	for this.serializedList.Len() > 0 {
		frontEl := this.serializedList.Remove(this.serializedList.Front()).(*util.TreeNode)
		length--
		if frontEl == nil {
			ans = ans + "nil,"
		} else {
			ans = ans + fmt.Sprintf("%d", frontEl.Val) + ","
			this.serializedList.PushBack(frontEl.Left)
			this.serializedList.PushBack(frontEl.Right)
		}
	}
	return ans[:len(ans)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *util.TreeNode {
	splitArr := strings.Split(data, ",")
	for _, val := range splitArr {
		if val == "nil" {
			this.deserializedList = append(this.deserializedList, nil)
		} else {
			value, err := strconv.Atoi(val)
			if err == nil {
				this.deserializedList = append(this.deserializedList, &util.TreeNode{Val: value})
			}
		}
	}
	start := 0
	for _, val := range this.deserializedList {
		if val != nil {
			left := 2*start + 1
			right := 2*start + 2
			if left < len(this.deserializedList) {
				val.Left = this.deserializedList[left]
			}
			if right < len(this.deserializedList) {
				val.Right = this.deserializedList[right]
			}
			start++
		}
	}
	if len(this.deserializedList) == 0 {
		return nil
	}
	return this.deserializedList[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	nums := []util.IntegerStruct{{1, true}, {2, true}, {3, true}, {-1, false}, {-1, false}, {4, true}, {5, true}}
	root := util.LevelOrder(nums)
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	fmt.Println(ans)
}

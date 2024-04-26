package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize_deserialize_bst(t *testing.T) {
	tests := []struct {
		nums []util.IntegerStruct
	}{
		{nums: []util.IntegerStruct{{2, true}, {1, true}, {3, true}}},
		{nums: []util.IntegerStruct{}},
	}
	for _, tt := range tests {
		root := util.LevelOrder(tt.nums)
		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		ans := deser.deserialize(data)
		assert.Equal(t, root, ans)
	}
}

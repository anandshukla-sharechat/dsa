package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serializationDeserialization(t *testing.T) {
	tests := []struct {
		nums []util.IntegerStruct
	}{
		{nums: []util.IntegerStruct{{1, true}, {2, true}, {3, true}, {-1, false}, {-1, false}, {4, true}, {5, true}}},
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

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_main(t *testing.T) {

	tt := struct {
		Input  []int
		Output []float64
	}{
		Input:  []int{1, 2, 3},
		Output: []float64{1, 1.5, 2},
	}
	medianObj := Constructor()

	for i, _ := range tt.Input {
		medianObj.AddNum(tt.Input[i])
		assert.Equal(t, tt.Output[i], medianObj.FindMedian())
	}
}

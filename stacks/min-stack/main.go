package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Arr []Struct
}

type Struct struct {
	Num int
	Min int
}

func Constructor() MinStack {
	return MinStack{
		Arr: make([]Struct, 0),
	}
}

func (this *MinStack) Push(val int) {
	topEl := math.MaxInt
	if len(this.Arr) > 0 {
		topEl = this.Arr[len(this.Arr)-1].Min
	}
	this.Arr = append(this.Arr, Struct{
		Num: val,
		Min: min(topEl, val),
	})
}

func (this *MinStack) Pop() {
	this.Arr = this.Arr[:len(this.Arr)-1]
}

func (this *MinStack) Top() int {
	return this.Arr[len(this.Arr)-1].Num
}

func (this *MinStack) GetMin() int {
	return this.Arr[len(this.Arr)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

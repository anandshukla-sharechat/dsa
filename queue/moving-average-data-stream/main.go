package main

import (
	"container/list"
	"fmt"
)

type MovingAverage struct {
	Queue  *list.List
	Window int
	Sum    int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Queue:  list.New(),
		Window: size,
		Sum:    0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.Queue.Len() == this.Window {
		this.Sum = this.Sum - this.Queue.Front().Value.(int)
		this.Queue.Remove(this.Queue.Front())
		this.Queue.PushBack(val)
		this.Sum = this.Sum + val
		return float64(this.Sum) / float64(this.Window)
	} else {
		this.Queue.PushBack(val)
		this.Sum = this.Sum + val
		return float64(this.Sum) / float64(this.Queue.Len())
	}
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func solve(size int, elements []int) []float64 {
	obj := Constructor(size)
	ans := make([]float64, 0)
	for _, val := range elements {
		ans = append(ans, obj.Next(val))
	}
	return ans
}

func main() {

	elements := []int{1, 10, 3, 5}
	ans := solve(3, elements)
	fmt.Println(ans)
}

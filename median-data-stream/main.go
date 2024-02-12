package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int
type MaxHeap []int

// Maxheap functions

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

// MinHeap functions

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

type MedianFinder struct {
	MaxPQ MaxHeap
	MinPQ MinHeap
}

func Constructor() MedianFinder {
	structPQ := MedianFinder{
		MaxPQ: make(MaxHeap, 0),
		MinPQ: make(MinHeap, 0),
	}
	heap.Init(&structPQ.MaxPQ)
	heap.Init(&structPQ.MinPQ)
	return structPQ
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.MaxPQ, num)
	numFromMaxPQ := heap.Pop(&this.MaxPQ).(int)
	heap.Push(&this.MinPQ, numFromMaxPQ)

	if len(this.MinPQ) > len(this.MaxPQ) {
		numFromMinPQ := heap.Pop(&this.MinPQ).(int)
		heap.Push(&this.MaxPQ, numFromMinPQ)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.MaxPQ) == len(this.MinPQ) {
		if len(this.MaxPQ) == 0 {
			return 0.0
		} else {
			median := (float64(this.MaxPQ[0]) + float64(this.MinPQ[0])) / 2.0
			return median
		}
	} else {
		return float64(this.MaxPQ[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// https://leetcode.com/problems/find-median-from-data-stream/description/
func main() {
	medianObj := Constructor()
	medianObj.AddNum(1)
	medianObj.AddNum(2)
	fmt.Println(medianObj.FindMedian())
	medianObj.AddNum(3)
	fmt.Println(medianObj.FindMedian())
}

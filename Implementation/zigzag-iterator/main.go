package main

import "fmt"

type ZigzagIterator struct {
	Arr1    []int
	Arr2    []int
	Index1  int
	Index2  int
	Pointer int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		Arr1:    v1,
		Arr2:    v2,
		Index1:  0,
		Index2:  0,
		Pointer: 1,
	}
}

func (this *ZigzagIterator) next() int {
	if this.Pointer == 1 {
		var temp int
		if this.Index1 < len(this.Arr1) {
			temp = this.Arr1[this.Index1]
			this.Index1++
			this.Pointer = 2
		} else {
			if this.Index2 < len(this.Arr2) {
				temp = this.Arr2[this.Index2]
				this.Index2++
				this.Pointer = 2
			}
		}
		return temp
	} else {
		var temp int
		if this.Index2 < len(this.Arr2) {
			temp = this.Arr2[this.Index2]
			this.Index2++
			this.Pointer = 1
		} else {
			if this.Index1 < len(this.Arr1) {
				temp = this.Arr1[this.Index1]
				this.Index1++
				this.Pointer = 1
			}
		}
		return temp
	}
}

func (this *ZigzagIterator) hasNext() bool {
	if this.Index1 < len(this.Arr1) || this.Index2 < len(this.Arr2) {
		return true
	} else {
		return false
	}
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */

func solve(a, b []int) []int {
	ans := make([]int, 0)
	obj := Constructor(a, b)
	for obj.hasNext() {
		ans = append(ans, obj.next())
	}
	return ans
}

func main() {
	a := []int{1, 2}
	b := []int{3, 4, 5, 6}
	fmt.Println(solve(a, b))
}

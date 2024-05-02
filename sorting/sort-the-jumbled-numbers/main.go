package main

import (
	"fmt"
	"sort"
)

type Arr struct {
	Num   int
	Val   int
	Index int
}

func sortJumbled(mapping []int, nums []int) []int {
	mp := make(map[int]int)
	for i, el := range mapping {
		mp[i] = el
	}
	ansArr := make([]Arr, 0)
	for idx, el := range nums {
		val := 0
		temp := make([]int, 0)
		p := 0
		for i := el; i != 0 || p >= 0; i = i / 10 {
			rem := i % 10
			encodedVal := mp[rem]
			temp = append(temp, encodedVal)
			p--
		}
		for j := len(temp) - 1; j >= 0; j-- {
			val = val*10 + temp[j]
		}
		ansArr = append(ansArr, Arr{
			Num:   el,
			Val:   val,
			Index: idx,
		})

	}
	sort.Slice(ansArr, func(i, j int) bool {
		if ansArr[i].Val == ansArr[j].Val {
			return ansArr[i].Index < ansArr[j].Index
		}
		return ansArr[i].Val < ansArr[j].Val
	})
	ans := make([]int, 0)
	for _, el := range ansArr {
		ans = append(ans, el.Num)
	}
	return ans
}

func main() {
	mapping, nums := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}
	ans := sortJumbled(mapping, nums)
	fmt.Println(ans)
}

package main

import (
	"container/list"
	"fmt"
)

func sequenceReconstruction(nums []int, sequences [][]int) bool {

	if len(nums) == 1 {
		if len(sequences) > 1 {
			return false
		} else {
			if nums[0] == sequences[0][0] {
				return true
			} else {
				return false
			}
		}
	}

	adj := make(map[int][]int)
	for _, val := range sequences {
		for i := 1; i < len(val); i++ {
			if _, ok := adj[val[i]]; !ok {
				adj[val[i]] = []int{val[i-1]}
			} else {
				adj[val[i]] = append(adj[val[i]], val[i-1])
			}
		}
	}
	degree := make(map[int]int)
	for key, _ := range adj {
		degree[key] = 0
	}
	for _, val := range adj {
		for _, v := range val {
			degree[v]++
		}
	}

	q := list.New()
	for index, val := range degree {
		if val == 0 {
			q.PushBack(index)
		}
	}
	flag := true
	subS := make([][]int, 0)
	for q.Len() > 0 {
		sz := q.Len()
		temp := make([]int, 0)
		for sz > 0 {
			el := q.Remove(q.Front()).(int)
			temp = append(temp, el)
			for _, val := range adj[el] {
				degree[val]--
				if degree[val] == 0 {
					q.PushBack(val)
				}
			}
			sz--
		}
		if len(temp) > 1 {
			flag = false
		}
		subS = append(subS, temp)
	}

	// prepare map for matching of numbers
	mp := make([]map[int]bool, 0)
	for i := len(subS) - 1; i >= 0; i-- {
		sub := subS[i]
		temp := make(map[int]bool)
		for _, val := range sub {
			temp[val] = true
		}
		mp = append(mp, temp)
	}

	i, count, k := 0, 0, 0
	for k < len(nums) && i < len(mp) {
		_, ok := mp[i][nums[k]]
		if ok {
			count++
			k++
		} else {
			i++
		}
	}

	if count == len(nums) {
		return flag
	} else {
		return false
	}
}

// https://leetcode.com/problems/sequence-reconstruction/description/
func main() {
	nums := []int{1, 2, 3}
	sequences := [][]int{{1, 2}, {1, 3}, {2, 3}}
	ans := sequenceReconstruction(nums, sequences)
	fmt.Println(ans)
}

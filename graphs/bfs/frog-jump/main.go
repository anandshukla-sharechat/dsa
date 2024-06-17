package main

import (
	"container/list"
	"fmt"
)

type jump struct {
	idx int
	k   int
}

func canCross(stones []int) bool {
	q := list.New()
	mp := make(map[[2]int]bool)
	q.PushBack(jump{idx: stones[0], k: 0})
	posMp := make(map[int]bool)
	for _, val := range stones {
		posMp[val] = true
	}

	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			el := q.Remove(q.Front()).(jump)
			if el.idx == stones[len(stones)-1] {
				return true
			}
			if el.k == 0 {
				if _, ok := posMp[el.idx+el.k+1]; ok {
					q.PushBack(jump{idx: el.idx + el.k + 1, k: el.k + 1})
				}
			} else {
				if _, ok := posMp[el.idx+el.k+1]; ok && !mp[[2]int{el.idx + el.k + 1, el.k + 1}] {
					q.PushBack(jump{idx: el.idx + el.k + 1, k: el.k + 1})
					mp[[2]int{el.idx + el.k + 1, el.k + 1}] = true
				}
				if _, ok := posMp[el.idx+el.k]; ok && !mp[[2]int{el.idx + el.k, el.k}] {
					q.PushBack(jump{idx: el.idx + el.k, k: el.k})
					mp[[2]int{el.idx + el.k, el.k}] = true
				}
				if _, ok := posMp[el.idx+el.k-1]; ok && !mp[[2]int{el.idx + el.k - 1, el.k - 1}] {
					q.PushBack(jump{idx: el.idx + el.k - 1, k: el.k - 1})
					mp[[2]int{el.idx + el.k - 1, el.k - 1}] = true
				}
			}
			mp[[2]int{el.idx, el.k}] = true
		}
	}

	return false

}

// HARD : https://leetcode.com/problems/frog-jump/
func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	ans := canCross(stones)
	fmt.Println(ans)
}

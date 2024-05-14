package main

import (
	"fmt"
	"sort"
)

func placeBalls(position []int, m int, safeDist int) bool {
	m--
	prev := position[0]
	for i := 1; i < len(position) && m != 0; i++ {
		if (position[i] - prev) >= safeDist {
			m--
			prev = position[i]
		}
	}
	if m == 0 {
		return true
	}
	return false
}

func maxDistance(position []int, m int) int {
	sort.Slice(position, func(i, j int) bool {
		return position[i] < position[j]
	})
	lower, upper := 0, int(1e9)
	ans := 0
	for lower <= upper {
		mid := lower + (upper-lower)/2
		if placeBalls(position, m, mid) {
			ans = mid
			lower = mid + 1
		} else {
			upper = mid - 1
		}
	}
	return ans
}

// https://leetcode.com/problems/magnetic-force-between-two-balls/description/
func main() {
	position, m := []int{79, 74, 57, 22}, 4
	ans := maxDistance(position, m)
	fmt.Println(ans)
}

package main

import (
	"fmt"
	"math"
)

func solve(p []int, h, mid int) bool {
	t := 0
	for _, val := range p {
		t += int(math.Ceil(float64(val) / float64(mid)))
	}
	if t <= h {
		return true
	} else {
		return false
	}
}

func minEatingSpeed(piles []int, h int) int {
	l, u, ans := 1, int(1e9), 0

	for l <= u {
		mid := l + (u-l)/2
		if solve(piles, h, mid) {
			ans = mid
			u = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

// https://leetcode.com/problems/koko-eating-bananas/description/
func main() {
	piles, h := []int{3, 6, 7, 11}, 8
	ans := minEatingSpeed(piles, h)
	fmt.Println(ans)
}

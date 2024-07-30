package main

import (
	"fmt"
	"math"
)

func solve(mid int64, cars int, ranks []int) bool {
	for _, val := range ranks {
		div := mid / int64(val)
		cars -= int(math.Sqrt(float64(div)))
		if cars <= 0 {
			break
		}
	}
	if cars <= 0 {
		return true
	} else {
		return false
	}
}

func repairCars(ranks []int, cars int) int64 {
	// binary search on minimizing value
	var left, right, ans, mid int64 = 0, int64(100 * cars * cars), 0, 0
	for left <= right {
		mid = left + (right-left)/2
		if solve(mid, cars, ranks) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// https://leetcode.com/problems/minimum-time-to-repair-cars/
func main() {
	ranks, cars := []int{4, 2, 3, 1}, 10
	ans := repairCars(ranks, cars)
	fmt.Println(ans)
}

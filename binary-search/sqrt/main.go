package main

import "fmt"

func mySqrt(x int) int {
	l, u, ans := 0, x, 0
	for l <= u {
		mid := (l + u) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			u = mid - 1
		}

	}
	return ans
}

// https://leetcode.com/problems/sqrtx/
func main() {
	x := 8
	ans := mySqrt(x)
	fmt.Println(ans)
}

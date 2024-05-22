package main

import "fmt"

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func bs(a, b []int, k int) bool {
	// mp := make(map[[k]int]int)
	for i := 0; i+k <= len(a); i++ {
		for j := 0; j+k <= len(b); j++ {
			if cmp(a[i:i+k], b[j:j+k]) {
				return true
			}
		}
	}

	return false
}

func findLength(nums1 []int, nums2 []int) int {
	l, u := 0, min(len(nums1), len(nums2))
	ans := 0
	for l <= u {
		mid := l + (u-l)/2
		if bs(nums1, nums2, mid) {
			ans = mid
			l = mid + 1
		} else {
			u = mid - 1
		}
	}
	return ans
}

// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
func main() {
	nums1, nums2 := []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}
	ans := findLength(nums1, nums2)
	fmt.Println(ans)
}

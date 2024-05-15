/*

Asked in DE SHAW first round

Given a string S and a string T, find the minimum window in S which will contain all the characters in T.

S = "ADOBECODEBANC"   -m
T = "ABC"             -n



Time complexity : O(n) + O(m) + O(m * log2(m))
Space complexity : O(m)
*/

package main

import (
	"fmt"
	"maps"
	"math"
)

func minWindow(s string, t string) string {
	// preprocess t
	mpT := make(map[int32]int)
	for _, el := range t {
		mpT[el]++
	}
	// preprocessing for s
	mpS := make(map[int]map[int32]int)
	for idx, el := range s {
		mpS[idx] = make(map[int32]int)
		if idx != 0 {
			maps.Copy(mpS[idx], mpS[idx-1])
		}
		mpS[idx][el]++
	}

	// final solution

	minL, ans, temp := math.MaxInt, "", ""
	for i := range s {

		// binary search
		lower, upper := 0, len(s)-1
		temp = ""
		for lower <= upper {
			mid := lower + (upper-lower)/2
			lastIdx := min(i+mid, len(s)-1)
			flag := -1
			for k, v := range mpT {
				if (mpS[lastIdx][k] - mpS[i-1][k]) < v {
					flag = 1
				}
			}
			if flag == -1 {
				temp = s[i:min(lastIdx+1, len(s))]
				upper = mid - 1
			} else {
				lower = mid + 1
			}
		}
		if minL > len(temp) && temp != "" {
			minL = len(temp)
			ans = temp
		}
	}
	return ans
}

/*
https://leetcode.com/problems/minimum-window-substring/
265 / 268 testcases passed, rest gave TLE
*/
func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t))
}

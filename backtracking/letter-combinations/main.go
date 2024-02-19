package main

import "fmt"

var LetterCombinationMap = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func solve(digits string, i int, res *[]string, sequence string) {
	if len(sequence) == len(digits) {
		*res = append(*res, sequence)
		return
	} else {
		key := LetterCombinationMap[rune(digits[i])]
		for j := 0; j < len(key); j++ {
			sequence = sequence + string(key[j])
			solve(digits, i+1, res, sequence)
			sequence = sequence[:len(sequence)-1]
		}
	}
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	solve(digits, 0, &res, "")
	return res
}

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func main() {
	digits := "23"
	res := letterCombinations(digits)
	for _, val := range res {
		fmt.Println(val)
	}
}

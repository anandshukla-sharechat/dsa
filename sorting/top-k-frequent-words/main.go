package main

import (
	"fmt"
	"sort"
)

type word struct {
	w    string
	freq int
}

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int)
	for _, el := range words {
		mp[el]++
	}
	pWords := make([]word, 0)
	for k, v := range mp {
		pWords = append(pWords, word{
			w:    k,
			freq: v,
		})
	}
	sort.Slice(pWords, func(i, j int) bool {
		if pWords[i].freq == pWords[j].freq {
			return pWords[i].w <= pWords[j].w
		} else {
			return pWords[i].freq > pWords[j].freq
		}
	})
	ans := make([]string, 0)
	for _, el := range pWords {
		ans = append(ans, el.w)
	}
	ans = ans[:min(k, len(ans))]
	return ans
}

func main() {
	words, k := []string{"i", "love", "leetcode", "i", "love", "coding"}, 2
	ans := topKFrequent(words, k)
	fmt.Println(ans)
}

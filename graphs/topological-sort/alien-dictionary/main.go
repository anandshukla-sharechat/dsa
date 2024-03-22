package main

import (
	"container/list"
	"fmt"
)

/*
There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.

You are given a list of strings words from the alien language's dictionary. Now it is claimed that the strings in words are
sorted lexicographically
 by the rules of this new language.

If this claim is incorrect, and the given arrangement of string in words cannot correspond to any order of letters, return "".

Otherwise, return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules.
If there are multiple solutions, return any of them.
*/

func alienOrder(words []string) string {
	adj := make(map[byte][]byte)
	degree := make(map[byte]int)

	for _, word := range words {
		for _, c := range []byte(word) {
			degree[c] = 0
			if _, ok := adj[c]; !ok {
				adj[c] = make([]byte, 0)
			}
		}
	}

	for i := 1; i < len(words); i++ {
		w1, w2 := words[i-1], words[i]
		j, k, n1, n2 := 0, 0, len(w1), len(w2)

		for j < n1 && k < n2 {
			if w1[j] != w2[k] {
				break
			}
			j++
			k++
		}
		if j == n1 {
			continue
		}
		if k == n2 {
			return ""
		}
		adj[w2[k]] = append(adj[w2[k]], w1[j])
		degree[w1[j]]++
	}
	q := list.New()
	for c, freq := range degree {
		if freq == 0 {
			q.PushBack(c)
		}
	}
	rev := ""
	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			el := q.Remove(q.Front()).(byte)
			rev = rev + string(el)
			for _, c := range adj[el] {
				degree[c]--
				if degree[c] == 0 {
					q.PushBack(c)
				}
			}
			sz--
		}
	}

	ans := ""
	for i := len(rev) - 1; i >= 0; i-- {
		ans = ans + string(rev[i])
	}

	if len(ans) == len(degree) {
		return ans
	} else {
		return ""
	}
}

func main() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	ans := alienOrder(words)
	fmt.Println(ans)
}

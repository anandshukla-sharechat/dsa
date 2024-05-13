package main

import "fmt"

func find(i string, mp map[string]string) string {
	if mp[i] == i {
		return i
	} else {
		res := find(mp[i], mp)
		mp[i] = res
		return res
	}
}

func union(i, j string, mp map[string]string) {
	res1 := find(i, mp)
	res2 := find(j, mp)
	if res1 != res2 {
		mp[res1] = mp[res2]
	}
}

func areSentencesSimilarTwo(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	mp := make(map[string]string)
	for _, val := range sentence1 {
		mp[val] = val
	}
	for _, val := range sentence2 {
		mp[val] = val
	}

	for _, val := range similarPairs {
		union(val[0], val[1], mp)
	}

	ans := make(map[string]int)

	for _, val := range mp {
		ans[find(val, mp)]++
	}

	if len(ans) == len(sentence1) {
		return true
	} else {
		return false
	}

}

/*
https://leetcode.com/problems/sentence-similarity-ii/

We can represent a sentence as an array of words, for example, the sentence "I am happy with leetcode" can be represented as arr = ["I","am",happy","with","leetcode"].

Given two sentences sentence1 and sentence2 each represented as a string array and given an array of string pairs similarPairs where similarPairs[i] = [xi, yi] indicates that the two words xi and yi are similar.

Return true if sentence1 and sentence2 are similar, or false if they are not similar.

Two sentences are similar if:

They have the same length (i.e., the same number of words)
sentence1[i] and sentence2[i] are similar.
Notice that a word is always similar to itself, also notice that the similarity relation is transitive. For example, if the words a and b are similar, and the words b and c are similar, then a and c are similar.
*/
func main() {
	sentence1, sentence2, similarPairs := []string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, [][]string{{"great", "good"}, {"fine", "good"}, {"drama", "acting"}, {"skills", "talent"}}
	ans := areSentencesSimilarTwo(sentence1, sentence2, similarPairs)
	fmt.Println(ans)
}

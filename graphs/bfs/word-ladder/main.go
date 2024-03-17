package main

import (
	"container/list"
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	mp := make(map[string]bool)
	for _, word := range wordList {
		mp[word] = true
	}
	q := list.New()
	q.PushBack(beginWord)
	count := 1
	for q.Len() > 0 {
		qSz := q.Len()
		for qSz > 0 {
			word := q.Remove(q.Front()).(string)
			qSz--
			delete(mp, word)
			if word == endWord {
				return count
			}
			for i := 0; i < len(word); i++ {
				for j := 0; j < 26; j++ {
					var char uint8 = uint8('a' + j)
					pre, post := word[:i], ""
					if (i + 1) < len(word) {
						post = word[i+1:]
					}
					newWord := pre + string(char) + post
					if _, ok := mp[newWord]; ok {
						q.PushBack(newWord)
					}
				}
			}
		}
		count++
	}
	return 0

}

/*
	https://leetcode.com/problems/word-ladder/description/
	Passed 50/51 cases, gave TLE for 51st case but ran individually within 1 second - both on local machine & leetcode
	Test case include in test file
	=== RUN   Test_ladderLength
	--- PASS: Test_ladderLength (0.27s)
	PASS

*/

func main() {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ans := ladderLength(beginWord, endWord, wordList)
	fmt.Println(ans)
}

package main

import (
	"container/list"
	heap_approach "first-unique-character-in-string/heap-approach"
	"fmt"
)

type Struct struct {
	Char  rune
	Index int
}

func firstUniqChar(s string) int {
	mp := make(map[rune]int64)
	queue := list.New()
	for i, val := range s {
		if _, ok := mp[val]; !ok {
			queue.PushBack(Struct{
				Char:  val,
				Index: i,
			})
		}
		mp[val]++
	}
	for el := queue.Front(); el != nil; el = el.Next() {
		if mp[el.Value.(Struct).Char] == 1 {
			return el.Value.(Struct).Index
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
	fmt.Println(heap_approach.FirstUniqChar(s))
}

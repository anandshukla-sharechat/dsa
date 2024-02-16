package heap_approach

import "container/heap"

type Struct struct {
	Val   rune
	Freq  int
	Index int
}

type StructArr []Struct

func (s StructArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StructArr) Less(i, j int) bool {
	if s[i].Freq == s[j].Freq {
		return s[i].Index < s[j].Index
	} else {
		return s[i].Freq < s[j].Freq
	}
}

func (s StructArr) Len() int {
	return len(s)
}

func (s *StructArr) Push(x any) {
	*s = append(*s, x.(Struct))
}

func (s *StructArr) Pop() any {
	old := *s
	oldLen := len(old)
	x := old[oldLen-1]
	*s = old[:oldLen-1]
	return x
}

func FirstUniqChar(s string) int {

	mp := make(map[rune]Struct)

	for i, char := range s {
		if val, ok := mp[char]; ok {
			val.Freq = val.Freq + 1
			mp[char] = val
		} else {
			mp[char] = Struct{
				Val:   char,
				Freq:  1,
				Index: i,
			}
		}
	}

	pq := make(StructArr, 0)
	heap.Init(&pq)
	for _, val := range mp {
		heap.Push(&pq, val)
	}

	if pq[0].Freq > 1 {
		return -1
	} else {
		return pq[0].Index
	}
}

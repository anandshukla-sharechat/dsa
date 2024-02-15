package stack

import "container/list"

func IsValid(s string) bool {
	st := list.New()

	mp := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if _, ok := mp[char]; ok {
			st.PushBack(char)
		} else {
			if st.Len() == 0 {
				return false
			}
			if mp[st.Back().Value.(rune)] != char {
				return false
			} else {
				st.Remove(st.Back())
			}
		}
	}
	return st.Len() == 0
}

package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	mystack "valid-parentheses/stack"
)

func isValid(s string) bool {
	st := stack.New()
	mp := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, char := range s {
		if _, ok := mp[char]; ok {
			st.Push(char)
		} else {
			topEl := st.Peek().(rune)
			matchingBracket := mp[topEl]
			if matchingBracket != char {
				return false
			} else {
				st.Pop()
			}
		}
	}
	return st.Len() == 0

}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(mystack.IsValid("()[]{}"))
}

package main

import "fmt"

func main() {
	var s string
	fmt.Println("Enter the string which you want to convert from square bracket to curl only for input")
	fmt.Scanf("%s", &s)
	ans := ""
	for _, val := range s {
		if val == '[' {
			ans = ans + "{"
		} else if val == ']' {
			ans = ans + "}"
		} else {
			ans = ans + string(val)
		}
	}
	fmt.Println(ans)
}

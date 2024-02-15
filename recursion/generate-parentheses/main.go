package main

import "fmt"

func solve(n, open, closed int, res *[]string, sequence string) {
	if open < n {
		sequence = sequence + "("
		open++
		solve(n, open, closed, res, sequence)
		open--
		sequence = sequence[:len(sequence)-1]
	}
	if closed < n && closed < open {
		sequence = sequence + ")"
		closed++
		solve(n, open, closed, res, sequence)
		closed--
		sequence = sequence[:len(sequence)-1]
	}
	if closed == n {
		*res = append(*res, sequence)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	open, closed := 0, 0
	sequence := ""
	solve(n, open, closed, &res, sequence)

	return res
}

func main() {
	var n int
	fmt.Println("Enter the value of n ")
	fmt.Scanf("%d", &n)
	res := generateParenthesis(n)
	for _, val := range res {
		fmt.Println(val)
	}
}

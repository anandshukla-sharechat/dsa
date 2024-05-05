package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func nthMagicalNumber(n int, a int, b int) int {
	mod := int(1e9 + 7)
	low, high := min(a, b), n*min(a, b)
	lcm := (a * b) / gcd(a, b)
	ans := -1
	for low <= high {
		mid := low + (high-low)/2
		rnk := mid/a + mid/b - mid/lcm
		if rnk >= n {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans % mod
}

func main() {
	n, a, b := 1000000000, 39999, 40000
	ans := nthMagicalNumber(n, a, b)
	fmt.Println(ans)
}

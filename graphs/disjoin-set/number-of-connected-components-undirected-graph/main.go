package main

import "fmt"

func find(i int, mp map[int]int) int {
	if mp[i] == i {
		return i
	} else {
		res := find(mp[i], mp)
		mp[i] = res
		return res
	}
}

func union(i, j int, mp map[int]int) {
	x, y := find(i, mp), find(j, mp)
	if x != y {
		mp[x] = y
	}
}

func countComponents(n int, edges [][]int) int {
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[i] = i
	}
	for _, val := range edges {
		union(val[0], val[1], mp)
	}
	countMap := make(map[int]int)
	for _, val := range mp {
		countMap[find(val, mp)]++
	}
	return len(countMap)
}

func main() {
	n, edges := 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	ans := countComponents(n, edges)
	fmt.Println(ans)
}

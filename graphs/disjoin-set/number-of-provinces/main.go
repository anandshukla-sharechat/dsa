package main

import "fmt"

func find(i int, parent *[]int) int {
	if (*parent)[i] == i {
		return i
	} else {
		result := find((*parent)[i], parent)
		(*parent)[i] = result
		return result
	}
}

func union(i, j int, parent *[]int) {
	a := find(i, parent)
	b := find(j, parent)
	if a != b {
		(*parent)[a] = b
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j, &arr)
			}
		}
	}
	mp := make(map[int]int)
	for _, val := range arr {
		mp[find(val, &arr)]++
	}
	return len(mp)
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	ans := findCircleNum(isConnected)
	fmt.Println(ans)
}

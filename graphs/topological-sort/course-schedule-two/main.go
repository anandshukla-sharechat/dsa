package main

import (
	"container/list"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {

	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}

	for _, val := range prerequisites {
		adj[val[0]] = append(adj[val[0]], val[1])
	}

	degree := make([]int, numCourses)
	for _, val := range adj {
		for _, el := range val {
			degree[el]++
		}
	}
	q := list.New()

	for i, val := range degree {
		if val == 0 {
			q.PushBack(i)
		}
	}

	ans := make([]int, 0)
	rev := make([]int, 0)

	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			el := q.Remove(q.Front()).(int)
			rev = append(rev, el)
			for i := 0; i < len(adj[el]); i++ {
				degree[adj[el][i]]--
				if degree[adj[el][i]] == 0 {
					q.PushBack(adj[el][i])
				}
			}
			sz--
		}
	}

	for i := len(rev) - 1; i >= 0; i-- {
		ans = append(ans, rev[i])
	}
	if len(ans) == numCourses {
		return ans
	} else {
		return []int{}
	}
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	ans := findOrder(numCourses, prerequisites)
	fmt.Println(ans)
}

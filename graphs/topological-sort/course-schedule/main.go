package main

import (
	"container/list"
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses+1)
	inDegree := make([]int, numCourses+1)
	for _, val := range prerequisites {
		if adj[val[0]] == nil {
			adj[val[0]] = []int{val[1]}
		} else {
			adj[val[0]] = append(adj[val[0]], val[1])
		}
	}
	for _, val := range adj {
		for _, preCourse := range val {
			inDegree[preCourse]++
		}
	}
	q := list.New()
	for i := 0; i <= numCourses; i++ {
		if inDegree[i] == 0 {
			q.PushBack(i)
		}
	}
	visited := make([]int, 0)
	for q.Len() > 0 {
		el := q.Remove(q.Front()).(int)
		visited = append(visited, el)
		for _, val := range adj[el] {
			inDegree[val]--
			if inDegree[val] == 0 {
				q.PushBack(val)
			}
		}
	}

	if len(visited) == (numCourses + 1) {
		return true
	} else {
		return false
	}
}

func main() {
	numCourses := 1
	prerequisites := [][]int{{1, 0}}
	ans := canFinish(numCourses, prerequisites)
	fmt.Println(ans)
}

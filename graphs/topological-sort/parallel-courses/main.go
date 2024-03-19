package main

import (
	"container/list"
	"fmt"
)

func minimumSemesters(n int, relations [][]int) int {
	adj := make(map[int][]int)
	degree := make(map[int]int)
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, 0)
		degree[i] = 0
	}

	for _, val := range relations {
		adj[val[1]] = append(adj[val[1]], val[0])
	}

	for _, val := range adj {
		for _, v := range val {
			degree[v]++
		}
	}

	q := list.New()
	for key, val := range degree {
		if val == 0 {
			q.PushBack(key)
		}
	}

	count := 0
	studied := 0
	for q.Len() > 0 {
		sz := q.Len()
		count++
		for sz > 0 {
			studied++
			el := q.Remove(q.Front()).(int)
			for _, val := range adj[el] {
				degree[val]--
				if degree[val] == 0 {
					q.PushBack(val)
				}
			}
			sz--
		}
	}
	if studied == n {
		return count
	} else {
		return -1
	}
}

// https://leetcode.com/problems/parallel-courses/description/
func main() {
	n, relations := 3, [][]int{{1, 3}, {2, 3}}
	ans := minimumSemesters(n, relations)
	fmt.Println(ans)
}

/*
You are given an integer n, which indicates that there are n courses labeled from 1 to n.
You are also given an array relations where relations[i] = [prevCoursei, nextCoursei], representing a prerequisite relationship between course prevCoursei and course nextCoursei: course prevCoursei has to be taken before course nextCoursei.

In one semester, you can take any number of courses as long as you have taken all the prerequisites in the previous semester for the courses you are taking.

Return the minimum number of semesters needed to take all courses. If there is no way to take all the courses, return -1
*/

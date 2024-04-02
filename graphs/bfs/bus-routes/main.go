package main

import (
	"container/list"
	"fmt"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	mp := make(map[int][]int)
	for bus, route := range routes {
		for _, stop := range route {
			if _, ok := mp[stop]; !ok {
				mp[stop] = []int{bus}
			} else {
				mp[stop] = append(mp[stop], bus)
			}
		}
	}
	q := list.New()
	visited := make(map[int]bool)
	for _, bus := range mp[source] {
		q.PushBack(bus)
		visited[bus] = true
	}

	ans := 1
	for q.Len() > 0 {
		sz := q.Len()
		for sz > 0 {
			bus := q.Remove(q.Front()).(int)
			for _, stop := range routes[bus] {
				if stop == target {
					return ans
				}
				for _, val := range mp[stop] {
					if !visited[val] {
						q.PushBack(val)
						visited[val] = true
					}
				}
			}
			sz--
		}
		ans++
	}
	return -1

}

/*
You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.

For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.
You will start at the bus stop source (You are not on any bus initially), and you want to go to the bus stop target. You can travel between bus stops by buses only.

Return the least number of buses you must take to travel from source to target. Return -1 if it is not possible.



Example 1:

Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
Output: 2
Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.
Example 2:

Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
Output: -1
*/

// https://leetcode.com/problems/bus-routes/description/
func main() {
	routes, source, target := [][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6
	ans := numBusesToDestination(routes, source, target)
	fmt.Println(ans)
}

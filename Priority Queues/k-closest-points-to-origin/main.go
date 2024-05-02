package main

import (
	"container/heap"
	"fmt"
)

type Coordinate struct {
	x int
	y int
	w int
}
type Maxheap []Coordinate

func (m Maxheap) Len() int {
	return len(m)
}

func (m Maxheap) Less(i, j int) bool {
	return m[i].w > m[j].w
}

func (m Maxheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *Maxheap) Push(x any) {
	*m = append(*m, x.(Coordinate))
}

func (m *Maxheap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	pq := make(Maxheap, 0)
	heap.Init(&pq)

	for _, el := range points {
		w := el[0]*el[0] + el[1]*el[1]
		heap.Push(&pq, Coordinate{
			x: el[0],
			y: el[1],
			w: w,
		})
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}
	ans := make([][]int, 0)
	for pq.Len() > 0 {
		c := heap.Pop(&pq).(Coordinate)
		ans = append(ans, []int{c.x, c.y})
	}
	return ans
}

// https://leetcode.com/problems/k-closest-points-to-origin/description/
func main() {
	points, k := [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2
	ans := kClosest(points, k)
	fmt.Println(ans)
}

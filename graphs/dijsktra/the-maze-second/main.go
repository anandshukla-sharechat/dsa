package main

import (
	"container/heap"
	"fmt"
)

type Item[T any] struct {
	value    T
	priority int
	index    int
}

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	queue := make(PriorityQueue[[]int], 0)
	heap.Push(&queue, &Item[[]int]{
		value:    start,
		priority: 0,
	})

	var j int
	var curDist, altDist, nextDist int
	for queue.Len() > 0 {
		curNode := heap.Pop(&queue).(*Item[[]int])
		curPosX, curPosY := curNode.value[0], curNode.value[1]

		curDist = maze[curPosX][curPosY] >> 1
		if curDist == 0 && (curPosX != start[0] || curPosY != start[1]) {
			curDist = 1<<63 - 1
		}

		if curPosX == destination[0] && curPosY == destination[1] {
			return curDist
		}

		if curDist < curNode.priority {
			continue
		}

		// up
		j = curPosX - 1
		for ; j >= 0 && maze[j][curPosY]&1 == 0; j-- {
		}
		j++
		nextDist = maze[j][curPosY] >> 1
		if nextDist == 0 && (j != start[0] || curPosY != start[1]) {
			nextDist = 1<<63 - 1
		}
		altDist = curDist + curPosX - j
		if altDist < nextDist {
			maze[j][curPosY] = altDist << 1
			heap.Push(&queue, &Item[[]int]{
				value:    []int{j, curPosY},
				priority: altDist,
			})
		}

		// down
		j = curPosX + 1
		for ; j < len(maze) && maze[j][curPosY]&1 == 0; j++ {
		}
		j--
		nextDist = maze[j][curPosY] >> 1
		if nextDist == 0 && (j != start[0] || curPosY != start[1]) {
			nextDist = 1<<63 - 1
		}
		altDist = curDist + j - curPosX
		if altDist < nextDist {
			maze[j][curPosY] = altDist << 1
			heap.Push(&queue, &Item[[]int]{
				value:    []int{j, curPosY},
				priority: altDist,
			})
		}

		// left
		j = curPosY - 1
		for ; j >= 0 && maze[curPosX][j]&1 == 0; j-- {
		}
		j++
		nextDist = maze[curPosX][j] >> 1
		if nextDist == 0 && (curPosX != start[0] || j != start[1]) {
			nextDist = 1<<63 - 1
		}
		altDist = curDist + curPosY - j
		if altDist < nextDist {
			maze[curPosX][j] = altDist << 1
			heap.Push(&queue, &Item[[]int]{
				value:    []int{curPosX, j},
				priority: altDist,
			})
		}

		// right
		j = curPosY + 1
		for ; j < len(maze[curPosX]) && maze[curPosX][j]&1 == 0; j++ {
		}
		j--
		nextDist = maze[curPosX][j] >> 1
		if nextDist == 0 && (curPosX != start[0] || j != start[1]) {
			nextDist = 1<<63 - 1
		}
		altDist = curDist + j - curPosY
		if altDist < nextDist {
			maze[curPosX][j] = altDist << 1
			heap.Push(&queue, &Item[[]int]{
				value:    []int{curPosX, j},
				priority: altDist,
			})
		}
	}

	return -1
}

func main() {
	maze, start, destination := [][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4}
	ans := shortestDistance(maze, start, destination)
	fmt.Println(ans)
}

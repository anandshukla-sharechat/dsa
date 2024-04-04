package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	oldLen := len(old)
	x := old[oldLen-1]
	*m = old[:oldLen-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	stops := 0
	tank := startFuel
	pq := make(MaxHeap, 0)
	heap.Init(&pq)
	prev := 0
	for _, val := range stations {
		stop := val[0]
		if stop >= target {
			stop = target
		}
		fuelSpent := tank - (stop - prev)
		for fuelSpent < 0 && pq.Len() > 0 {
			fuelSpent += heap.Pop(&pq).(int)
			stops++
		}
		if fuelSpent < 0 {
			return -1
		}
		if target == stop {
			return stops
		}
		tank = fuelSpent
		prev = stop
		heap.Push(&pq, val[1])
	}
	fuelSpent := tank - (target - prev)
	for fuelSpent < 0 && pq.Len() > 0 {
		fuelSpent += heap.Pop(&pq).(int)
		stops++
	}
	if fuelSpent >= 0 {
		return stops
	} else {
		return -1
	}
}

func main() {
	target, startFuel, stations := 100, 50, [][]int{{25, 25}, {50, 50}}
	ans := minRefuelStops(target, startFuel, stations)
	fmt.Println(ans)
}

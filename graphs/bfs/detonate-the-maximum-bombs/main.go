package main

import (
	"container/list"
	"fmt"
	"math"
)

type Coordinates struct {
	X int
	Y int
	R int
}

func square(a int) int {
	return a * a
}

func LiesWithinCircle(x1 int, y1 int, x2 int, y2 int, radius int) bool {
	if (square(x1-x2) + square(y1-y2)) <= square(radius) {
		return true
	} else {
		return false
	}
}

func maximumDetonation(bombs [][]int) int {

	queue := list.New()
	n := len(bombs)
	// for duplicates
	dup := 0
	dupMap := make(map[[2]int]int)
	for i := 0; i < n; i++ {
		dupMap[[2]int{bombs[i][0], bombs[i][1]}]++
	}
	for _, val := range dupMap {
		if val >= 2 {
			dup += (val - 1)
		}
	}
	res := math.MinInt
	for i := 0; i < n; i++ {
		vis := make(map[[2]int]bool)
		queue.PushBack(Coordinates{X: bombs[i][0], Y: bombs[i][1], R: bombs[i][2]})
		for queue.Len() > 0 {
			source := queue.Front().Value.(Coordinates)
			vis[[2]int{source.X, source.Y}] = true
			queue.Remove(queue.Front())
			for j := 0; j < n; j++ {
				if i != j && (LiesWithinCircle(source.X, source.Y, bombs[j][0], bombs[j][1], source.R) || LiesWithinCircle(source.X, source.Y, bombs[j][0], bombs[j][1], bombs[j][2])) && !vis[[2]int{bombs[j][0], bombs[j][1]}] {
					queue.PushBack(Coordinates{X: bombs[j][0], Y: bombs[j][1], R: bombs[j][2]})
				}
			}
		}
		res = max(res, len(vis))
	}
	return res + dup
}

func main() {
	bombs := [][]int{{4, 4, 3}, {4, 4, 3}}
	fmt.Println(maximumDetonation(bombs))
}

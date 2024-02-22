package main

import (
	"container/list"
	"fmt"
	"math"
)

type Coordinate struct {
	X int
	Y int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	score := make([][]int, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			temp[j] = math.MaxInt
		}
		score[i] = temp
	}
	queue := list.New()

	direction := []Coordinate{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				score[i][j] = 0
				queue.PushBack(Coordinate{X: i, Y: j})
				for queue.Len() > 0 {
					source := queue.Front().Value.(Coordinate)
					queue.Remove(queue.Front())
					for _, dir := range direction {
						x := source.X + dir.X
						y := source.Y + dir.Y
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && score[x][y] > score[source.X][source.Y]+1 {
							score[x][y] = min(score[x][y], score[source.X][source.Y]+1)
							queue.PushBack(Coordinate{X: x, Y: y})
						}
					}
				}

			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if score[i][j] == math.MaxInt {
					return -1
				}
				ans = max(ans, score[i][j])
			}
		}
	}
	return ans
}

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}

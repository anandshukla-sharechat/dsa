package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

type Weight struct {
	Weight int
	Source int
}

type WeightArr []Weight

func (w WeightArr) Len() int {
	return len(w)
}

func (w WeightArr) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WeightArr) Less(i, j int) bool {
	return w[i].Weight > w[j].Weight
}

func (w *WeightArr) Push(x any) {
	*w = append(*w, x.(Weight))
}

func (w *WeightArr) Pop() any {
	old := *w
	oldLen := len(old)
	x := old[oldLen-1]
	*w = old[:oldLen-1]
	return x
}

func maxResult(nums []int, k int) int {
	ans := 0
	pqMax := make(WeightArr, 0)
	n := len(nums)
	heap.Init(&pqMax)
	heap.Push(&pqMax, Weight{
		Weight: nums[0],
		Source: 0,
	})
	mark := make([]int, n)
	for i := 0; i < n; i++ {
		mark[i] = -1
	}

	for len(pqMax) > 0 {
		topEl := heap.Pop(&pqMax).(Weight)
		i := topEl.Source + 1
		upper_limit := topEl.Source + k
		for ; i <= min(upper_limit, n-1); i++ {
			if mark[i] < topEl.Weight+nums[i] {
				heap.Push(&pqMax, Weight{
					Weight: topEl.Weight + nums[i],
					Source: i,
				})
				if i == n-1 {
					ans = max(ans, topEl.Weight+nums[i])
				}
				mark[i] = topEl.Weight + nums[i]
			}

		}
	}
	return ans
}

func solveRecursion(nums []int, k int, ans *int, i int, weight int, n int) {

	for j := i + 1; j <= min(i+k, n-1); j++ {
		weight = weight + nums[j]
		if j == (n - 1) {
			*ans = max(*ans, weight)
		}
		solveRecursion(nums, k, ans, j, weight, n)
		weight = weight - nums[j]
	}
}

func solveByRecursion(nums []int, k int) int {
	ans := math.MinInt
	solveRecursion(nums, k, &ans, 0, nums[0], len(nums))
	return ans
}

// solveByMonotonicQueue : Maintain a monotonic decreasing queue & remove elements when a increased score is encountered
func solveByMonotonicQueue(nums []int, k int) int {
	n := len(nums)
	queue := list.New()
	score := make([]int, n)
	score[0] = nums[0]
	queue.PushBack(0)

	for i := 1; i < n; i++ {
		// remove older elements
		for queue.Len() > 0 && queue.Front().Value.(int) < (i-k) {
			queue.Remove(queue.Front())
		}
		score[i] = score[queue.Front().Value.(int)] + nums[i]
		for queue.Len() > 0 && score[queue.Back().Value.(int)] <= score[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(i)
	}
	return score[n-1]
}

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 2
	fmt.Println(maxResult(nums, k))
	fmt.Println(solveByRecursion(nums, k))
	fmt.Println(solveByMonotonicQueue(nums, k))
}

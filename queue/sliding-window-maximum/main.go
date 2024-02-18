package main

import (
	"container/list"
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	queue := list.New()

	for i := 0; i < len(nums); i++ {
		for queue.Len() > 0 && queue.Front().Value.(int) <= (i-k) {
			queue.Remove(queue.Front())
		}
		for queue.Len() > 0 && nums[queue.Back().Value.(int)] <= nums[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(i)
		if i >= k-1 {
			ans = append(ans, nums[queue.Front().Value.(int)])
		}
	}
	return ans
}

func main() {
	nums, k := []int{1, 3, -1, -3, 5, 3, 6, 7}, 3
	fmt.Println(maxSlidingWindow(nums, k))
}

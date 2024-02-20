package main

import (
	"container/list"
	"fmt"
)

type Store struct {
	Key int
	Val int
}

type LRUCache struct {
	Queue    *list.List
	Map      map[int]*list.Element
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Queue:    list.New(),
		Map:      make(map[int]*list.Element),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Map[key]; ok {
		// update key to last of queue
		this.Queue.MoveToBack(val)
		return val.Value.(Store).Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// if key is already present
	if val, ok := this.Map[key]; ok {
		val.Value = Store{Key: key, Val: value}
		this.Queue.MoveToBack(val)
		return
	} else {
		if this.Queue.Len() == this.Capacity {
			delete(this.Map, this.Queue.Front().Value.(Store).Key)
			this.Queue.Remove(this.Queue.Front())
		}
		el := this.Queue.PushBack(Store{Key: key, Val: value})
		this.Map[key] = el
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // cache is {1=1}
	lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4

}

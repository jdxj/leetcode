package s622

import "math"

type MyCircularQueue struct {
	cap        uint64
	start, end uint64

	cache []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap:   uint64(k),
		start: 0,
		end:   0,
		cache: make([]int, k, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (mcq *MyCircularQueue) EnQueue(value int) bool {
	amount := mcq.end - mcq.start
	// 满了
	if amount == mcq.cap {
		return false
	}
	// 序号达到最大
	if mcq.end == math.MaxUint64 {
		mcq.start = mcq.start % mcq.cap
		mcq.end = mcq.start + amount
	}
	// 正常情况
	mcq.cache[mcq.end%mcq.cap] = value
	mcq.end++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (mcq *MyCircularQueue) DeQueue() bool {
	// 没有元素
	if mcq.IsEmpty() {
		return false
	}
	mcq.start++
	return true
}

/** Get the front item from the queue. */
func (mcq *MyCircularQueue) Front() int {
	// 没有元素
	if mcq.IsEmpty() {
		return -1
	}
	return mcq.cache[mcq.start%mcq.cap]
}

/** Get the last item from the queue. */
func (mcq *MyCircularQueue) Rear() int {
	amount := mcq.end - mcq.start
	if amount == 0 {
		return -1
	}
	lastIndex := mcq.start + amount - 1
	return mcq.cache[lastIndex%mcq.cap]
}

/** Checks whether the circular queue is empty or not. */
func (mcq *MyCircularQueue) IsEmpty() bool {
	return mcq.end-mcq.start == 0
}

/** Checks whether the circular queue is full or not. */
func (mcq *MyCircularQueue) IsFull() bool {
	return mcq.end-mcq.start == mcq.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

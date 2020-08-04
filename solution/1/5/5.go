// 155
package s15

import "math"

var (
	stackBuf    = make([]int, 500)
	minStaskBuf = make([]int, 500)
)

type MinStack struct {
	stackLen int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stackBuf = stackBuf[:0]
	minStaskBuf = minStaskBuf[:0]

	ms := MinStack{}
	return ms
}

func (this *MinStack) Push(x int) {
	minTop := this.GetMin()
	minStaskBuf = append(minStaskBuf, min(x, minTop))
	stackBuf = append(stackBuf, x)
	this.stackLen++
}

func (this *MinStack) Pop() {
	if this.stackLen == 0 {
		return
	}
	stackBuf = stackBuf[:this.stackLen-1]
	minStaskBuf = minStaskBuf[:this.stackLen-1]
	this.stackLen--
}

func (this *MinStack) Top() int {
	return stackBuf[this.stackLen-1]
}

func (this *MinStack) GetMin() int {
	if this.stackLen == 0 {
		return math.MaxInt64
	}
	return minStaskBuf[this.stackLen-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func min(p1, p2 int) int {
	if p1 > p2 {
		p1 = p2
	}
	return p1
}

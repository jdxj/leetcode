// 384
package s38

import (
	"math/rand"
	"time"
)

var (
	buf    = make([]int, 1000)
	origin []int
)

type Solution struct {
	r       *rand.Rand
	numsLen int
}

func Constructor(nums []int) Solution {
	buf = buf[:0]
	for i := 0; i < len(nums); i++ {
		buf = append(buf, nums[i])
	}
	origin = nums
	sol := Solution{
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		numsLen: len(nums),
	}
	return sol
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := this.numsLen - 1; i >= 0; i-- {
		idx := this.r.Intn(i + 1)
		buf[i], buf[idx] = buf[idx], buf[i]
	}
	return buf
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

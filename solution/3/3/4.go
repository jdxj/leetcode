// 334
package s33

import "math"

func increasingTriplet(nums []int) bool {
	s, m := math.MaxInt64, math.MaxInt64
	for _, v := range nums {
		if v <= s {
			s = v
		} else if v <= m {
			m = v
		} else {
			return true
		}
	}
	return false
}

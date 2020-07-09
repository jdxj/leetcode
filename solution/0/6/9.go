package s69

func mySqrtFloat64(x float64) float64 {
	lower, upper := float64(0), x
	for {
		mid := lower + (upper-lower)/2
		// 误差计算
		if mid*mid > x && mid*mid-x < 0.001 {
			return mid
		}
		if mid*mid < x && x-mid*mid < 0.001 {
			return mid
		}
		// 大了
		if mid > x/mid {
			upper = mid
		} else {
			lower = mid
		}
	}
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	res := mySqrtFloat64(float64(x))
	return int(res)
}

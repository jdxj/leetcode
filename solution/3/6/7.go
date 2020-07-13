// 367
package s36

// 牛顿迭代
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}

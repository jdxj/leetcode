// 202
package s20

func isHappy(n int) bool {
	record := make(map[int]struct{})

	for n = next(n); ; n = next(n) {
		// 如果结果是1, 说明是快乐数
		if n == 1 {
			return true
		}
		// 在记录中查找, 如果找到了说明不是快乐数
		if _, ok := record[n]; ok {
			return false
		}
		record[n] = struct{}{}
	}
}

func next(n int) int {
	var result int
	for ; n > 0; n = n / 10 {
		remainder := n % 10
		result += remainder * remainder
	}
	return result
}

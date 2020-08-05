// 204
package s20

func countPrimes(n int) int {
	isPrim := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrim[i] = true
	}
	// sqrt(n) ~ n 由内部更改值
	for i := 2; i*i < n; i++ {
		if !isPrim[i] {
			continue
		}
		// 筛
		for j := i * i; j < n; j += i {
			isPrim[j] = false
		}
	}

	var count int
	for i := 2; i < n; i++ {
		if isPrim[i] {
			count++
		}
	}
	return count
}

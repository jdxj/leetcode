// 461
package s46

func hammingDistance(x int, y int) int {
	return hammingWeight(x ^ y)
}

func hammingWeight(num int) int {
	var count int
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

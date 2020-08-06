// 191
package s19

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

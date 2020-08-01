// 119
package s11

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for j := 0; j <= rowIndex; j++ {
		result[j] = combination(rowIndex, j)
	}
	return result
}

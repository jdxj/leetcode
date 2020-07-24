// 498
package s49

func findDiagonalOrder(matrix [][]int) []int {
	rowLen := len(matrix)
	if rowLen == 0 {
		return nil
	}
	colLen := len(matrix[0])

	if rowLen == 1 {
		return matrix[0]
	}
	if colLen == 1 {
		result := make([]int, 0, rowLen)
		for i := 0; i < rowLen; i++ {
			result = append(result, matrix[i][0])
		}
		return result
	}

	result := make([]int, 0, rowLen*colLen)

	var i, j int
	for z := 0; z <= rowLen+colLen-2; z++ {
		if z%2 != 0 {
			i = 0 // 明确 i 初始为0
			for z-i >= colLen {
				i++
			}
			for j = z - i; i < rowLen && j >= 0 && j < colLen; {
				result = append(result, matrix[i][j])
				i++
				j = z - i
			}
		} else {
			j = 0 // 明确 j 初始为0
			for z-j >= rowLen {
				j++
			}
			for i = z - j; j < colLen && i >= 0 && i < rowLen; {
				result = append(result, matrix[i][j])
				j++
				i = z - j
			}
		}
	}
	return result
}

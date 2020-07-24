package interview_questions

// 这题的关键点是用当前坐标 [i, j] 推导旋转坐标
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]                    // 当前位置 0
			matrix[i][j] = matrix[n-j-1][i]         // 前一位置 1
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1] // 位置1的前一位置 2
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1] // 位置0 的后一位置
			matrix[j][n-i-1] = temp
		}
	}
}

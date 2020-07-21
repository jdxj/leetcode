// 36
package s03

func isValidSudoku(board [][]byte) bool {
	return helper1(board) && helper2(board) && helper3(board)
}

// 检测行
func helper1(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		row := make(map[byte]struct{})
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := row[board[i][j]]; ok {
				return false
			} else {
				row[board[i][j]] = struct{}{}
			}
		}
	}
	return true
}

// 检测列
func helper2(board [][]byte) bool {
	for j := 0; j < len(board[0]); j++ {
		col := make(map[byte]struct{})
		for i := 0; i < len(board); i++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := col[board[i][j]]; ok {
				return false
			} else {
				col[board[i][j]] = struct{}{}
			}
		}
	}
	return true
}

// 检测块
func helper3(board [][]byte) bool {
	blk := make([]map[byte]struct{}, 9)
	for i := range blk {
		blk[i] = make(map[byte]struct{})
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := blk[i/3*3+j/3][board[i][j]]; ok {
				return false
			} else {
				blk[i/3*3+j/3][board[i][j]] = struct{}{}
			}
		}
	}
	return true
}

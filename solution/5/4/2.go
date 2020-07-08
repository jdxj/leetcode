package s542

import (
	"container/list"
)

// 输入:
//   0 0 0
//   0 1 0
//   1 1 1
// 输出:
//   0 0 0
//   0 1 0
//   1 2 1
type kv struct {
	i, j int
}

// 非常耗时
func updateMatrix(matrix [][]int) [][]int {
	queue := list.New()
	record := make(map[int]struct{})

	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				queue.PushBack(kv{i, j})
				record[i*cols+j] = struct{}{}
			}
		}
	}

	for queue.Front() != nil {
		pos := queue.Remove(queue.Front()).(kv)
		// 上下左右 各位置+1
		if pos.i-1 >= 0 {
			if _, ok := record[(pos.i-1)*cols+pos.j]; !ok {
				matrix[pos.i-1][pos.j] = matrix[pos.i][pos.j] + 1
				record[(pos.i-1)*cols+pos.j] = struct{}{}
				queue.PushBack(kv{pos.i - 1, pos.j})
			}
		}
		if pos.i+1 < rows {
			if _, ok := record[(pos.i+1)*cols+pos.j]; !ok {
				matrix[pos.i+1][pos.j] = matrix[pos.i][pos.j] + 1
				record[(pos.i+1)*cols+pos.j] = struct{}{}
				queue.PushBack(kv{pos.i + 1, pos.j})
			}
		}
		if pos.j-1 >= 0 {
			if _, ok := record[pos.i*cols+pos.j-1]; !ok {
				matrix[pos.i][pos.j-1] = matrix[pos.i][pos.j] + 1
				record[pos.i*cols+pos.j-1] = struct{}{}
				queue.PushBack(kv{pos.i, pos.j - 1})
			}
		}
		if pos.j+1 < cols {
			if _, ok := record[pos.i*cols+pos.j+1]; !ok {
				matrix[pos.i][pos.j+1] = matrix[pos.i][pos.j] + 1
				record[pos.i*cols+pos.j+1] = struct{}{}
				queue.PushBack(kv{pos.i, pos.j + 1})
			}
		}
	}
	return matrix
}

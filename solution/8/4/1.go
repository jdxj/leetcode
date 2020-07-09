package s841

// [[1],[2],[3],[]]
func canVisitAllRooms(rooms [][]int) bool {
	rows := len(rooms)
	if rows == 0 {
		return true
	}

	stack := make([]int, 0)
	record := make(map[int]struct{})

	record[0] = struct{}{}             // 记录0号房间已访问过
	stack = append(stack, rooms[0]...) // 将0号房间中的钥匙入栈

	for len(stack) != 0 {
		// 弹出一个钥匙
		key := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果 key 号房间已访问过, 则跳过
		if _, ok := record[key]; ok {
			continue
		}
		record[key] = struct{}{}             // 记录 key 号房间已访问过
		stack = append(stack, rooms[key]...) // 将 key 号房间中的钥匙入栈
	}

	for i := range rooms {
		// 出现每打开过的房间
		if _, ok := record[i]; !ok {
			return false
		}
	}
	return true
}

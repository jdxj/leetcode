package s739

// T := []int{73, 74, 75, 71, 69, 72, 76, 73}
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	stack := []int{0} // stack 存储了某个温度 x 在 T 中的索引 i, 即: T[stack[top]] = x

	for i := 1; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] { // 如果栈顶代表的温度小于当前温度
			preI := stack[len(stack)-1]
			result[preI] = i - preI
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

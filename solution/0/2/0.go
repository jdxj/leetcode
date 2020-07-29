// 20
package s02

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := make([]string, 0)

	for _, c := range s {
		if _, ok := m[string(c)]; !ok { // 如果是左括号, 那么入栈
			stack = append(stack, string(c))
		} else if len(stack) == 0 || m[string(c)] != stack[len(stack)-1] { // 如果是右括号, 那么比较栈顶元素
			// 如果栈中没有元素, 或者栈顶元素不是对应的左括号, 则返回 false
			return false
		} else { // 与栈顶括号匹配了, 弹出左括号
			stack = stack[:len(stack)-1]
		}
	}
	// 检查是否有剩余
	return len(stack) == 0
}

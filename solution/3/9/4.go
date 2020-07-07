package s394

import "strconv"

// s = "3[a2[c]]"
func decodeString(s string) string {
	var stack []string
	for _, c := range s {
		if c != ']' {
			// push
			stack = append(stack, string(c))
			continue
		}

		var sub string
		for stack[len(stack)-1] < "0" || stack[len(stack)-1] > "9" {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top == "[" {
				break
			}
			sub = top + sub
		}

		var kStr string
		for stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
			kStr = stack[len(stack)-1] + kStr
			// pop
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
		k, err := strconv.Atoi(kStr)
		if err != nil {
			panic(err)
		}
		for i := 0; i < k; i++ {
			stack = append(stack, sub)
		}
	}

	var result string
	for _, v := range stack {
		result += v
	}
	return result
}

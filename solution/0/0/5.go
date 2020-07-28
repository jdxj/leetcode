// 5
package s00

func longestPalindrome(s string) string {
	amount := len(s)
	if amount < 2 {
		return s
	}

	dp := make([][]bool, amount)
	for i := range dp {
		dp[i] = make([]bool, amount)
	}

	var result string
	// 枚举字串长度, l+1 是真正的长度, 因为 l 是从0开始的
	for l := 0; l < amount; l++ {
		// 枚举 i,j 状态
		for i := 0; i < amount; i++ {
			j := i + l
			if j >= amount {
				break
			}
			if l == 0 { // 此时 i == j, 说明起始是一个字符串时
				dp[i][j] = true
			} else if l == 1 { // 此时 i == j-1, 说明起始是两个字符
				dp[i][j] = s[i] == s[j]
			} else { // 其它情况则推断
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			// 记录最大值
			if dp[i][j] && l+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

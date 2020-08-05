// 13
package s01

var (
	nums = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	preNum, sum := nums[s[0]], 0
	for i := 1; i < sLen; i++ {
		if preNum < nums[s[i]] {
			preNum = -preNum
		}
		sum += preNum
		preNum = nums[s[i]]
	}
	return sum + preNum
}

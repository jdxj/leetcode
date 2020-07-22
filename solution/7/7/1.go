// 771
package s77

func numJewelsInStones(J string, S string) int {
	JMap := make(map[byte]struct{})
	for i := 0; i < len(J); i++ {
		JMap[J[i]] = struct{}{}
	}

	var result int
	for i := 0; i < len(S); i++ {
		if _, ok := JMap[S[i]]; ok {
			result++
		}
	}
	return result
}

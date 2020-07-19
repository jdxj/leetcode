// 205
package s20

func isIsomorphic(s string, t string) bool {
	conS := make(map[uint8]int)
	conT := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := conS[s[i]]; !ok {
			conS[s[i]] = i
		}
		if _, ok := conT[t[i]]; !ok {
			conT[t[i]] = i
		}
		if conS[s[i]] != conT[t[i]] {
			return false
		}
	}
	return true
}

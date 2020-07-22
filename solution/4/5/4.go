// 454
package s45

func fourSumCount(A []int, B []int, C []int, D []int) int {
	lenA, lenB := len(A), len(B)
	lenC, lenD := len(C), len(D)

	// kind: amount
	sumAB := make(map[int]int, lenA*lenB)
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			sumAB[A[i]+B[j]]++
		}
	}

	var result int
	for i := 0; i < lenC; i++ {
		for j := 0; j < lenD; j++ {
			if amount, ok := sumAB[-(C[i] + D[j])]; ok {
				result += amount
			}
		}
	}
	return result
}

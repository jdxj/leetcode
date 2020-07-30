// 118
package s11

import (
	"math/big"
)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			row[j] = combination(i, j)
		}
		result[i] = row
	}
	return result
}

var (
	factorialCache = make(map[int64]*big.Int, 25)
)

func factorial(n int64) *big.Int {
	if v, ok := factorialCache[n]; ok {
		return v
	}

	if n == 0 {
		factorialCache[0] = big.NewInt(1)
		return factorialCache[0]
	}

	bigN := big.NewInt(n)
	bigN.Mul(bigN, factorial(n-1))
	factorialCache[n] = bigN
	return bigN
}

func combination(n, k int) int {
	nF := factorial(int64(n))
	kF := factorial(int64(k))
	nkF := factorial(int64(n - k))
	// cp
	nF = new(big.Int).SetBytes(nF.Bytes())
	kF = new(big.Int).SetBytes(kF.Bytes())
	nkF = new(big.Int).SetBytes(nkF.Bytes())
	return int(nF.Div(nF, kF.Mul(kF, nkF)).Int64())
}

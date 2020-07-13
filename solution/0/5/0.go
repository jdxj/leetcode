// 50
package s05

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 { // 如果 n 小于0, 那么先计算分母: -n
		return 1 / myPow(x, -n)
	}
	if n%2 != 0 { // 如果 n 是奇数, 那么变成偶数
		return x * myPow(x, n-1)
	}
	// 第一个参数可以理解为: x 已乘了多少次, 变化是已乘了: 2次, 4次, 8次
	// 第二个参数可以解释为还需翻倍次数
	return myPow(x*x, n/2)
	// 2*2   4
	// 4*4   2
	// 16*16 1
	// 16*16 0
}

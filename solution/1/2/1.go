// 121
package s12

import "math"

func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt64, 0
	for _, v := range prices {
		if v < minPrice {
			// 更新最新最低价格的原因是:
			//   如果未来价格呈不变或上升, 那么最新最低价格有优势,
			//   如果未来价格呈下降, 那么最新价格与最新最低价格的差也是优势.
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}

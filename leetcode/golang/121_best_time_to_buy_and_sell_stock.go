package leetcode

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

// 动态规划，dp[i]= max{dp[i-1], prices[i]-prices[min]}
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	min := prices[0]
	dp[0] = 0
	for i := 1; i < n; i++ {
		price := prices[i]
		if price < min {
			min = price
		}
		dpi := price - min
		if dpi > dp[i-1] {
			dp[i] = dpi
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

// 动态规划改进，优化空间复杂度
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 0; i < n; i++ {
		price := prices[i]
		if price < min {
			min = price
		}
		temp := price - min
		if temp > max {
			max = temp
		}
	}
	return max
}

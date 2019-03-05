package leetcode

// https://leetcode-cn.com/problems/coin-change/
/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1

示例 2:
输入: coins = [2], amount = 3
输出: -1

说明:
你可以认为每种硬币的数量是无限的。
*/

/**
思路：动态规划，
状态：d(i) = j,表示兑换元需要最少j个银币，
状态转移方程：d(i) = min{d(i-vj) + 1}, 其中i-vj>=0, vj表示第j枚银币的面值
*/

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	amountArrs := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		amountArrs[i] = math.MaxInt32
	}
	amountArrs[0] = 0
	for i := 1; i <= amount; i++ {
		minCount := amountArrs[i]
		for _, coin := range coins {
			if i >= coin {
				temp := amountArrs[i-coin] + 1
				if temp < minCount {
					minCount = temp
				}
			}
		}
		amountArrs[i] = minCount
	}
	if amountArrs[amount] == math.MaxInt32 {
		return -1
	}
	return amountArrs[amount]
}

// 稍微优化一下，减少一点不必要的大小判断
func coinChange2(coins []int, amount int) int {
	n := amount + 1
	a := make([]int, n)
	for i := 1; i < n; i++ {
		a[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for i := coin; i < n; i++ {
			if a[i] > a[i-coin]+1 {
				a[i] = a[i-coin] + 1
			}
		}
	}

	if a[amount] == math.MaxInt32 {
		return -1
	}
	return a[amount]
}

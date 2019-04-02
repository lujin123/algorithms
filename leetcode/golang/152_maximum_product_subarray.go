package leetcode

import (
	"math"
)

// https://leetcode-cn.com/problems/maximum-product-subarray/
/**
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

func maxProduct(nums []int) int {
	max := 1
	min := 1
	ans := math.MinInt64
	for _, num := range nums {
		if num < 0 {
			max, min = min, max
		}
		max *= num
		if max < num {
			max = num
		}
		min *= num
		if min > num {
			min = num
		}
		if max > ans {
			ans = max
		}
	}
	return ans
}

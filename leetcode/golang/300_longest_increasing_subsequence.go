package leetcode

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:
可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。

进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

/**
思路：
1. 在数组中保存这个升序的值，之后遇到小值一直往数组前面搜索，找到比他还要小的值为止，之后替换
如果想要优化到O(n log n)可以使用二分查找，这样查询时间就是O(log n)了，还是灰常精妙的
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	var stack []int
	stack = append(stack, nums[0])
	for i := 1; i < n; i++ {
		lastIndex := len(stack) - 1
		if stack[lastIndex] >= nums[i] {
			for lastIndex--; lastIndex >= 0 && stack[lastIndex] >= nums[i]; lastIndex-- {
			}
			lastIndex++
			stack[lastIndex] = nums[i]
		} else {
			stack = append(stack, nums[i])
		}
	}
	return len(stack)
}

// 2. 动态规划
func lengthOfLIS2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	max := 1
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

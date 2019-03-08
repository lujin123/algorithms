package leetcode

// https://leetcode-cn.com/problems/product-of-array-except-self/

/**
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/

/**
思路：先用一个数组，从左往右算出当前位左边的乘积，保存在当前位，之后从后遍历原始数组，累积邮编的乘积，之后在和当前位左边的积相乘即可
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	res := 1
	for i := n - 1; i >= 0; i-- {
		left[i] *= res
		res *= nums[i]
	}

	return left
}

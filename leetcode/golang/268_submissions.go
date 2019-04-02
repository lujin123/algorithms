package leetcode

// https://leetcode-cn.com/problems/missing-number/submissions/

/**
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:

输入: [3,0,1]
输出: 2
示例 2:

输入: [9,6,4,2,3,5,7,0,1]
输出: 8
说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
*/
// 直接求和公式算出结果，在算出数组和，相减就是结果
func missingNumber(nums []int) int {
	n := len(nums)
	total := (1 + n) * n / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

// 直接异或求值，两个相同的值异或后会相互抵消，剩下的就是结果值
func missingNumber2(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res ^= v
		res ^= i
	}
	return res
}

// 直接存到对应的数组中，遍历数组查看缺少的值
func missingNumber3(nums []int) int {
	array := make([]int, len(nums)+1)
	for _, v := range nums {
		array[v] = v
	}
	ans := 0
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == 0 {
			ans = i
			break
		}
	}
	return ans
}

package leetcode

// https://leetcode-cn.com/problems/rotate-array/
/**
189. 旋转数组

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的原地算法。
*/

// 1. 将前k个数和后k个元素交互，这样前K个元素就是最终的位置
// 之后类推，k个元素之后的元素当做一个新数组，再交换即可
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n == 0 || k == 0 {
		return
	}
	start := 0
	for n > 0 && k != 0 {
		for i := 0; i < k; i++ {
			nums[i+start], nums[n-k+i+start] = nums[n-k+i+start], nums[i+start]
		}
		start += k
		n -= k
		k %= n
	}
}

// 2. 将数组前n-k个元素进行翻转，后k个元素也进行翻转，之后全部元素再翻转即可
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n == 0 || k == 0 {
		return
	}
	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

package leetcode

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
/**
33. 搜索旋转排序数组

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4

示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

// 思路：主要是判断哪个部分是有序的
// 1. 如果中间值比最后的值小，则右边有序，
// 2. 如果中间值比最开始的值大，则左边有序
// 之后再继续二分搜索即可
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		num := nums[m]
		if target == num {
			return m
		}
		if num < nums[r] {
			if target == nums[r] {
				return r
			} else if target > num && target < nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target == nums[l] {
				return l
			} else if target < num && target > nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}

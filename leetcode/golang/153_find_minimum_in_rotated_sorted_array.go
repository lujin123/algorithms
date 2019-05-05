package leetcode

import "math"

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
/**
153. 寻找旋转排序数组中的最小值

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
*/
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func findMin2(nums []int) int {
	min := math.MaxInt32
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		temp := nums[m]
		if nums[l] < temp {
			temp = nums[l]
		}
		if nums[r] < temp {
			temp = nums[r]
		}
		if temp < min {
			min = temp
		}
		if nums[m] > nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return min
}

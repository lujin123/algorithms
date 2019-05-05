package leetcode

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
/**
81. 搜索旋转排序数组 II

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
进阶:

这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/

// 思路：
// 重复的元素会导致中间的元素和前后的数字相同，导致没办法判断出单调递增的部分，
// 所以将开头和结尾重复的元素去掉即可，如果开头和结尾元素相同，结尾元素也直接去掉
func searchII(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		for l+1 <= r && nums[l] == nums[l+1] {
			l++
		}
		for r-1 >= l && (nums[r] == nums[r-1] || nums[l] == nums[r]) {
			r--
		}

		m := (l + r) / 2
		num := nums[m]
		if target == num {
			return true
		}

		if num < nums[r] {
			if target == nums[r] {
				return true
			} else if target > num && target < nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target == nums[l] {
				return true
			} else if target < num && target > nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return false
}

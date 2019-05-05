package leetcode

// https://leetcode-cn.com/problems/first-missing-positive/
/**
41. 缺失的第一个正数

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
*/

// 思路：主要是交换元素位置，将数组元素值放到对应的下标中，一遍循环之后，再次遍历数组，
// 查找到数组中下标和元素大小不相同的第一个就是缺失的值，或者遍历结束都没找到，就是数组大小加一的值
// 注意一点，需要排除掉大于数组大小的元素，因为这些元素不可能是缺失的第一个元素，
// 如果数组中有数缺失，那就会导致前面的数组无法被填满，随意缺失的第一个元素一定会在原数组中找到，
// 如果找不到说明不缺元素，那缺失的第一个就是数组的大小+1
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		num := nums[i]
		if num > 0 && num <= n && nums[num-1] != nums[i] {
			nums[i], nums[num-1] = nums[num-1], nums[i]
		} else {
			i++
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

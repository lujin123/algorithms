package leetcode

// https://leetcode-cn.com/problems/find-the-duplicate-number/
/**
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2
示例 2:

输入: [3,1,3,4,2]
输出: 3

说明：
	1. 不能更改原数组（假设数组是只读的）。
	2. 只能使用额外的 O(1) 的空间。
	3. 时间复杂度小于 O(n^2) 。
	4. 数组中只有一个重复的数字，但它可能不止重复出现一次。
*/

// 1. 二分查找+鸽笼原理，时间复杂度 = O(n*log n)
func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		middle := left + ((right - left) >> 1)
		var sum int
		for i := range nums {
			if nums[i] <= middle {
				sum++
			}
		}
		if sum > middle {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

// 2. 利用快慢指针，查找循环链表一样
func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	finder := 0
	for {
		slow = nums[slow]
		finder = nums[finder]
		if slow == finder {
			return finder
		}
	}
}

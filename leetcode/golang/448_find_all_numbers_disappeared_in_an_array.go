package leetcode

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
/**
448. 找到所有数组中消失的数字

给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, n] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:

输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]
*/

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; {
		if i+1 != nums[i] && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	var ans []int
	for i := range nums {
		if i+1 != nums[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}

package leetcode

// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/
/**
442. 数组中重复的数据

给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。

找到所有出现两次的元素。

你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

示例：

输入:
[4,3,2,7,8,2,3,1]

输出:
[2,3]
*/

// 在输入数组中用数字的正负来表示该位置所对应数字是否已经出现过。
// 遍历输入数组，给对应位置的数字取相反数，如果已经是负数，说明前面已经出现过，直接放入输出数组。
// 注意这题元素只会重复两次，如果重复的次数多了就不好用了
func findDuplicates(nums []int) []int {
	n := len(nums)
	var ans []int
	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		} else {
			ans = append(ans, num)
		}
	}
	return ans
}

func findDuplicates2(nums []int) []int {
	n := len(nums)
	var ans []int
	for i := 0; i < n; {
		if nums[i] > 0 && i+1 != nums[i] {
			if nums[i] != nums[nums[i]-1] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
				continue
			}
			ans = append(ans, nums[i])
			nums[i] = 0
		}
		i++
	}
	return ans
}

package leetcode

// https://leetcode-cn.com/problems/majority-element/
/**
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
*/

// 思路：直接多用空间，保存每个数的出现次数，很暴力
func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	middle := len(nums) / 2
	for _, v := range nums {
		if val, ok := numMap[v]; ok {
			res := val + 1
			if res > middle {
				return v
			}
			numMap[v] = res
		} else {
			numMap[v] = 1
		}
	}
	return 0
}

// 计数器如果是0，则将当前的数赋给结果值，之后遇到相同的计数器加一，不同的减一，减到零换下个数
func majorityElement2(nums []int) int {
	res, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if res == num {
			count++
		} else {
			count--
		}
	}
	return res
}

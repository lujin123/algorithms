package leetcode

import (
	"sort"
)

// https://leetcode-cn.com/problems/subsets-ii/
/**
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{[]int{}}
	}
	sort.Ints(nums)

	var results [][]int
	lastLen := 0
	for i := 0; i < n; i++ {
		num := nums[i]
		var (
			temp [][]int
			left int
		)
		flag := i > 0 && num == nums[i-1]
		if flag {
			left = len(results) - lastLen
		}

		for j := left; j < len(results); j++ {
			v := results[j]
			a := make([]int, len(v))
			copy(a, v)
			temp = append(temp, a)
		}
		lastLen = len(temp)

		for _, v := range temp {
			results = append(results, append(v, num))
		}

		if !flag {
			results = append(results, []int{num})
			lastLen += 1
		}
	}

	results = append(results, []int{})
	return results
}

// 精简代码
func subsetsWithDup2(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	size := 0
	for i, num := range nums {
		start := 0

		if i > 0 && num == nums[i-1] {
			start = size
		}
		size = len(res)
		for j := start; j < size; j++ {
			n := len(res[j])
			dummy := make([]int, n+1)
			copy(dummy, res[j])
			dummy[n] = num
			res = append(res, dummy)
		}
	}
	return res
}

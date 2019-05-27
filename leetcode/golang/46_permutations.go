package leetcode

// https://leetcode-cn.com/problems/permutations/
/**
46. 全排列

给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n < 1 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{nums}
	}

	allrange(nums, 0, n-1, &res)
	return res
}

func allrange(nums []int, start, end int, res *[][]int) {
	if start == end {
		newNums := append([]int{}, nums...)
		*res = append(*res, newNums)
	} else {
		for i := start; i <= end; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			allrange(nums, start+1, end, res)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

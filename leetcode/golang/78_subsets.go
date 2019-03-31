package leetcode

// https://leetcode-cn.com/problems/subsets/

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

// 这个解法很有意思，直接从左到右遍历，遇到一个新元素，就将原来的子集添加这个新元素组成
// 新的子集添加到数组中，遍历结束额外增加一个空数组即可，简直精妙
func subsets(nums []int) [][]int {
	var results [][]int
	for _, num := range nums {
		var temp [][]int
		for _, v := range results {
			a := make([]int, len(v))
			copy(a, v)
			temp = append(temp, a)
		}
		for _, v := range temp {
			results = append(results, append(v, num))
		}
		results = append(results, []int{num})
	}
	results = append(results, []int{})
	return results
}

//使用二进制表示， nums数组长度为1的个数，例如n=3, 那么从000~111，计算每个1所在位置，
// 映射到nums数组下标中即可
func subsets2(nums []int) [][]int {
	var results [][]int

	n := 1 << uint(len(nums))
	results = append(results, []int{})

	for i := 1; i < n; i++ {
		k := 0
		var res []int
		for j := i; j > 0; j >>= 1 {
			if (j & 1) == 1 {
				res = append(res, nums[k])
			}
			k++
		}
		results = append(results, res)
	}

	return results
}

// 精简下代码
func subsets3(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			nn := len(res[i])
			a := make([]int, nn+1)
			copy(a, res[i])
			a[nn] = num
			res = append(res, a)
		}
	}
	return res
}

package leetcode

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
/**
15. 三数之和

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := n - 1
			if nums[i]+nums[l]+nums[l+1] > 0 || nums[i]+nums[r]+nums[r-1] < 0 {
				continue
			}
			for l < r {
				a := nums[i] + nums[l] + nums[r]
				if a == 0 {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}

				} else if a < 0 {
					l++
				} else {
					r--
				}
			}
		}
	}
	if len(res) == 0 {
		return [][]int{}
	}
	return res
}

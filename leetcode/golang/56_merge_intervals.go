package leetcode

import (
	"sort"
)

// https://leetcode-cn.com/problems/merge-intervals/
/**
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

/**
type Interval struct {
	Start int
	End   int
}
*/
type Interval struct {
	Start int
	End   int
}

// 先按照start值从小到大排序，之后合并重叠的部分即可,时间复杂度应该是是`O(nlogn)`，主要看排序算法了
func merge(intervals []Interval) []Interval {
	n := len(intervals)
	if n == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var res []Interval
	index := 1
	interval := intervals[0]
	for index < n {
		next := intervals[index]
		if interval.End >= next.Start {
			if interval.End < next.End {
				interval.End = next.End
			}
		} else {
			res = append(res, interval)
			interval = next
		}
		index++
	}
	res = append(res, interval)
	return res
}

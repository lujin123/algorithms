package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{1, 10}}, withMerge([][]int{[]int{2, 3}, []int{4, 5}, []int{6, 7}, []int{8, 9}, []int{1, 10}}))
	assert.ElementsMatch(t, [][]int{[]int{0, 4}}, withMerge([][]int{[]int{1, 4}, []int{0, 4}}))
	assert.ElementsMatch(t, [][]int{[]int{0, 5}}, withMerge([][]int{[]int{1, 5}, []int{0, 4}}))
	assert.ElementsMatch(t, [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}}, withMerge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	assert.ElementsMatch(t, [][]int{[]int{1, 5}}, withMerge([][]int{[]int{1, 4}, []int{4, 5}}))

	assert.ElementsMatch(t, [][]int{[]int{1, 2}, []int{4, 5}}, withMerge([][]int{[]int{1, 2}, []int{4, 5}}))
}

func withMerge(nums [][]int) [][]int {
	var intervals []Interval
	for _, num := range nums {
		intervals = append(intervals, Interval{Start: num[0], End: num[1]})
	}
	results := merge(intervals)
	var ans [][]int
	for _, result := range results {
		ans = append(ans, []int{result.Start, result.End})
	}
	return ans
}

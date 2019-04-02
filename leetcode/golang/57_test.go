package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInterval(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{5, 7}}, withInsert([][]int{}, []int{5, 7}))
	assert.ElementsMatch(t, [][]int{[]int{1, 10}}, withInsert([][]int{[]int{2, 3}, []int{4, 5}, []int{6, 7}, []int{8, 9}}, []int{1, 10}))
	assert.ElementsMatch(t, [][]int{[]int{1, 5}, []int{6, 9}}, withInsert([][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}))
	assert.ElementsMatch(t, [][]int{[]int{1, 4}}, withInsert([][]int{[]int{1, 1}, []int{2, 4}}, []int{1, 2}))
	assert.ElementsMatch(t, [][]int{[]int{1, 2}, []int{3, 10}, []int{12, 16}}, withInsert([][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}, []int{4, 8}))
	assert.ElementsMatch(t, [][]int{[]int{1, 2}, []int{3, 5}}, withInsert([][]int{[]int{1, 2}, []int{4, 5}}, []int{3, 5}))

	assert.ElementsMatch(t, [][]int{[]int{1, 2}, []int{4, 5}}, withInsert([][]int{[]int{1, 2}}, []int{4, 5}))
}

func withInsert(nums [][]int, newInterval []int) [][]int {
	var intervals []Interval
	for _, num := range nums {
		intervals = append(intervals, Interval{Start: num[0], End: num[1]})
	}
	results := insert(intervals, Interval{Start: newInterval[0], End: newInterval[1]})
	var ans [][]int
	for _, result := range results {
		ans = append(ans, []int{result.Start, result.End})
	}
	return ans
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetsWithDup(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{2}, []int{2, 2}, []int{2, 2, 2}, []int{2, 2, 2, 2}, []int{2, 2, 2, 2, 2}, []int{}}, subsetsWithDup([]int{2, 2, 2, 2, 2}))
	assert.ElementsMatch(t, [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{1, 3}, []int{1, 2, 3}, []int{2, 3}, []int{3}, []int{}}, subsetsWithDup([]int{1, 2, 3}))

	assert.ElementsMatch(t, [][]int{[]int{0}, []int{0, 3}, []int{3}, []int{0, 5}, []int{0, 3, 5}, []int{3, 5}, []int{5}, []int{0, 7}, []int{0, 3, 7}, []int{3, 7}, []int{0, 5, 7}, []int{0, 3, 5, 7}, []int{3, 5, 7}, []int{5, 7}, []int{7}, []int{0, 9}, []int{0, 3, 9}, []int{3, 9}, []int{0, 5, 9}, []int{0, 3, 5, 9}, []int{3, 5, 9}, []int{5, 9}, []int{0, 7, 9}, []int{0, 3, 7, 9}, []int{3, 7, 9}, []int{0, 5, 7, 9}, []int{0, 3, 5, 7, 9}, []int{3, 5, 7, 9}, []int{5, 7, 9}, []int{7, 9}, []int{9}, []int{}}, subsetsWithDup([]int{0, 3, 5, 7, 9}))

	assert.ElementsMatch(t, [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{1, 2, 2}, []int{2, 2}, []int{}}, subsetsWithDup([]int{1, 2, 2}))

	assert.ElementsMatch(t, [][]int{[]int{}}, subsetsWithDup([]int{}))
	assert.ElementsMatch(t, [][]int{[]int{2}, []int{2, 2}, []int{}}, subsetsWithDup([]int{2, 2}))
}

func TestSubsetsWithDup2(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{2}, []int{2, 2}, []int{2, 2, 2}, []int{2, 2, 2, 2}, []int{2, 2, 2, 2, 2}, []int{}}, subsetsWithDup2([]int{2, 2, 2, 2, 2}))
	assert.ElementsMatch(t, [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{1, 3}, []int{1, 2, 3}, []int{2, 3}, []int{3}, []int{}}, subsetsWithDup2([]int{1, 2, 3}))

	assert.ElementsMatch(t, [][]int{[]int{0}, []int{0, 3}, []int{3}, []int{0, 5}, []int{0, 3, 5}, []int{3, 5}, []int{5}, []int{0, 7}, []int{0, 3, 7}, []int{3, 7}, []int{0, 5, 7}, []int{0, 3, 5, 7}, []int{3, 5, 7}, []int{5, 7}, []int{7}, []int{0, 9}, []int{0, 3, 9}, []int{3, 9}, []int{0, 5, 9}, []int{0, 3, 5, 9}, []int{3, 5, 9}, []int{5, 9}, []int{0, 7, 9}, []int{0, 3, 7, 9}, []int{3, 7, 9}, []int{0, 5, 7, 9}, []int{0, 3, 5, 7, 9}, []int{3, 5, 7, 9}, []int{5, 7, 9}, []int{7, 9}, []int{9}, []int{}}, subsetsWithDup2([]int{0, 3, 5, 7, 9}))

	assert.ElementsMatch(t, [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{1, 2, 2}, []int{2, 2}, []int{}}, subsetsWithDup2([]int{1, 2, 2}))

	assert.ElementsMatch(t, [][]int{[]int{}}, subsetsWithDup2([]int{}))
	assert.ElementsMatch(t, [][]int{[]int{2}, []int{2, 2}, []int{}}, subsetsWithDup2([]int{2, 2}))
}

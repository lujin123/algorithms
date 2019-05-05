package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	assert.ElementsMatch(t, []int{2, 3}, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func TestFindDuplicates2(t *testing.T) {
	assert.ElementsMatch(t, []int{2, 3}, findDuplicates2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	assert.ElementsMatch(t, []int{1, 2, 3, 7}, findDuplicates2([]int{1, 3, 2, 7, 7, 2, 3, 1}))
}

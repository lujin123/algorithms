package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	assert.Equal(t, 6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	assert.Equal(t, 3, lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 1, lengthOfLIS([]int{1}))
	assert.Equal(t, 0, lengthOfLIS([]int{}))
	assert.Equal(t, 1, lengthOfLIS([]int{2, 2}))
	assert.Equal(t, 3, lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	assert.Equal(t, 5, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 9, 100}))
}

func TestLengthOfLIS2(t *testing.T) {
	assert.Equal(t, 6, lengthOfLIS2([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	assert.Equal(t, 3, lengthOfLIS2([]int{10, 9, 2, 5, 3, 4}))
	assert.Equal(t, 4, lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 1, lengthOfLIS2([]int{1}))
	assert.Equal(t, 0, lengthOfLIS2([]int{}))
	assert.Equal(t, 1, lengthOfLIS2([]int{2, 2}))
	assert.Equal(t, 3, lengthOfLIS2([]int{4, 10, 4, 3, 8, 9}))
	assert.Equal(t, 5, lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 9, 100}))
}

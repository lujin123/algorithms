package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 0, rob([]int{}))
	assert.Equal(t, 1, rob([]int{1}))
	assert.Equal(t, 2, rob([]int{1, 2}))
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
	assert.Equal(t, 4, rob([]int{2, 1, 1, 2}))
}

func TestRob2(t *testing.T) {
	assert.Equal(t, 0, rob2([]int{}))
	assert.Equal(t, 1, rob2([]int{1}))
	assert.Equal(t, 2, rob2([]int{1, 2}))
	assert.Equal(t, 4, rob2([]int{1, 2, 3, 1}))
	assert.Equal(t, 12, rob2([]int{2, 7, 9, 3, 1}))
	assert.Equal(t, 4, rob2([]int{2, 1, 1, 2}))
}

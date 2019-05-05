package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchII(t *testing.T) {
	assert.Equal(t, true, searchII([]int{3, 1, 2, 2, 2}, 1))
	assert.Equal(t, true, searchII([]int{1, 3, 1, 1, 1}, 3))
	assert.Equal(t, true, searchII([]int{1, 1, 1, 3, 1}, 3))
	assert.Equal(t, true, searchII([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.Equal(t, false, searchII([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
	assert.Equal(t, 1, search([]int{4, 5, 1, 2, 3}, 5))
	assert.Equal(t, 2, search([]int{5, 1, 3}, 3))
	assert.Equal(t, 2, search([]int{4, 5, 0, 1, 2, 3}, 0))
	assert.Equal(t, 0, search([]int{5, 1, 3}, 5))
	assert.Equal(t, 0, search([]int{4}, 4))
	assert.Equal(t, 2, search([]int{4, 5, 6, 7, 0, 1, 2}, 6))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	assert.Equal(t, -1, search([]int{4}, 3))
	assert.Equal(t, -1, search([]int{}, 3))
}

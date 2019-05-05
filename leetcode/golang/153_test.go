package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{4, 5, 6, 7, 8, 1, 2, 3}))
	assert.Equal(t, 1, findMin([]int{4, 5, 1, 2, 3}))
	assert.Equal(t, 1, findMin([]int{5, 1, 3}))
	assert.Equal(t, 0, findMin([]int{4, 5, 0, 1, 2, 3}))
	assert.Equal(t, 1, findMin([]int{5, 1, 3}))
	assert.Equal(t, 4, findMin([]int{4}))
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 4, findMin([]int{4}))
	assert.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func TestFindMin2(t *testing.T) {
	assert.Equal(t, 1, findMin2([]int{4, 5, 6, 7, 8, 1, 2, 3}))
	assert.Equal(t, 1, findMin2([]int{4, 5, 1, 2, 3}))
	assert.Equal(t, 1, findMin2([]int{5, 1, 3}))
	assert.Equal(t, 0, findMin2([]int{4, 5, 0, 1, 2, 3}))
	assert.Equal(t, 1, findMin2([]int{5, 1, 3}))
	assert.Equal(t, 4, findMin2([]int{4}))
	assert.Equal(t, 0, findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 0, findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 0, findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 4, findMin2([]int{4}))
	assert.Equal(t, 1, findMin2([]int{3, 4, 5, 1, 2}))
	assert.Equal(t, 0, findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
}

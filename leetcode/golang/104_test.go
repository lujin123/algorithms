package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	assert.Equal(t, 0, maxDepth(nil))
	assert.Equal(t, 3, maxDepth(createFulllikeTreeWithArray([]int{3, 9, 20, 0, 0, 15, 7})))
}

func TestMaxDepth2(t *testing.T) {
	assert.Equal(t, 0, maxDepth2(nil))
	assert.Equal(t, 3, maxDepth2(createFulllikeTreeWithArray([]int{3, 9, 20, 0, 0, 15, 7})))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	assert.Equal(t, true, isBalanced(nil))
	assert.Equal(t, true, isBalanced(createFulllikeTreeWithArray([]int{3, 9, 20, 0, 0, 15, 7})))
	assert.Equal(t, false, isBalanced(createFulllikeTreeWithArray([]int{1, 2, 2, 3, 3, 0, 0, 4, 4})))

}

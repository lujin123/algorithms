package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{1, 2, 2, 3, 4, 4, 3})))

	assert.Equal(t, false, isSymmetric(createFulllikeTreeWithArray([]int{1, 2, 2, 0, 3, 0, 3})))

	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{1})))
	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{})))
}

func TestIsSymmetric2(t *testing.T) {
	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{1, 2, 2, 3, 4, 4, 3})))

	assert.Equal(t, false, isSymmetric2(createFulllikeTreeWithArray([]int{1, 2, 2, 0, 3, 0, 3})))

	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{1})))
	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{})))
}

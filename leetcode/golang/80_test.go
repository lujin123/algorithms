package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesII(t *testing.T) {
	assert.Equal(t, 5, removeDuplicatesII([]int{1, 1, 1, 2, 2, 3}))
	assert.Equal(t, 7, removeDuplicatesII([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func TestRemoveDuplicatesII2(t *testing.T) {
	assert.Equal(t, 5, removeDuplicatesII2([]int{1, 1, 1, 2, 2, 3}))
	assert.Equal(t, 7, removeDuplicatesII2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5}, calcListNode([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{1, 2, 3, 4}, calcListNode([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal(t, []int{2, 3, 4, 5}, calcListNode([]int{1, 2, 3, 4, 5}, 5))
}

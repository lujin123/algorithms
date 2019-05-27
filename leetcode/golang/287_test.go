package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
	assert.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
	assert.Equal(t, 2, findDuplicate([]int{1, 2, 2, 2, 5, 6, 7}))
}

func TestFindDuplicate2(t *testing.T) {
	assert.Equal(t, 2, findDuplicate2([]int{1, 3, 4, 2, 2}))
	assert.Equal(t, 3, findDuplicate2([]int{3, 1, 3, 4, 2}))
	assert.Equal(t, 2, findDuplicate2([]int{1, 2, 2, 2, 5, 6, 7}))
}

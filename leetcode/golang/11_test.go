package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func TestMaxArea2(t *testing.T) {
	assert.Equal(t, 49, maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	assert.EqualValues(t, 0, maxSubArray([]int{0}))
	assert.EqualValues(t, -1, maxSubArray([]int{-1, -2, -3}))
	assert.EqualValues(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

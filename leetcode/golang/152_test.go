package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	assert.EqualValues(t, 6, maxProduct([]int{2, 3, -2, 4}))
	assert.EqualValues(t, 0, maxProduct([]int{-2, 0, -1}))
}

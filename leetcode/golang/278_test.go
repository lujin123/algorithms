package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))
}

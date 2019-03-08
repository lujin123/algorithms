package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasGroupsSizeX(t *testing.T) {
	assert.Equal(t, true, hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	assert.Equal(t, false, hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	assert.Equal(t, false, hasGroupsSizeX([]int{1}))
	assert.Equal(t, false, hasGroupsSizeX([]int{1, 2}))
	assert.Equal(t, true, hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
}

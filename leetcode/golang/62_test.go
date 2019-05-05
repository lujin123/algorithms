package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 1, uniquePaths(1, 1))
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 28, uniquePaths(7, 3))
}

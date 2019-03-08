package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	assert.Equal(t, -1, minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}}))
	assert.Equal(t, 11, minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

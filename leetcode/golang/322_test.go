package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))
	assert.Equal(t, -1, coinChange([]int{2}, 3))
}

func TestCoinChange2(t *testing.T) {
	assert.Equal(t, 3, coinChange2([]int{1, 2, 5}, 11))
	assert.Equal(t, -1, coinChange2([]int{2}, 3))
}

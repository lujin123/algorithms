package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitII(t *testing.T) {
	assert.Equal(t, 7, maxProfitII([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, maxProfitII([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfitII([]int{7, 6, 4, 3, 1}))
}

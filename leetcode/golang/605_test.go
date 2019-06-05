package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	// 下面这个测试用例错了，违反了种植规则
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 1, 1, 0, 0}, 2))
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 0, 1, 0, 0}, 2))
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	assert.Equal(t, false, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

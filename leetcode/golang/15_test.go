package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.ElementsMatch(t, [][]int{}, threeSum([]int{-1, 0, 3, 50, -1, -4}))
}

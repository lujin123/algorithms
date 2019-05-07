package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}, generateMatrix(3))
	assert.Equal(t, [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}, generateMatrix(4))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 2, 1}, []int{3, 1, 2}}, permute([]int{1, 2, 3}))
}

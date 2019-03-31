package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, []int{7, 0, 8}, execAddTwoNumbers([]int{2, 4, 3}, []int{5, 6, 4}))
	assert.Equal(t, []int{0, 7, 4}, execAddTwoNumbers([]int{5}, []int{5, 6, 4}))
	assert.Equal(t, []int{5}, execAddTwoNumbers([]int{5}, []int{}))
	assert.Equal(t, []int{8, 0, 1}, execAddTwoNumbers([]int{9}, []int{9, 9}))
}

func TestAddTwoNumbers2(t *testing.T) {
	assert.Equal(t, []int{7, 0, 8}, execAddTwoNumbers2([]int{2, 4, 3}, []int{5, 6, 4}))
	assert.Equal(t, []int{0, 7, 4}, execAddTwoNumbers2([]int{5}, []int{5, 6, 4}))
	assert.Equal(t, []int{5}, execAddTwoNumbers2([]int{5}, []int{}))
	assert.Equal(t, []int{8, 0, 1}, execAddTwoNumbers2([]int{9}, []int{9, 9}))
}

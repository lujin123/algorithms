package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	assert.Equal(t, []int{4, 5, 1, 2, 3}, calcRotateRight([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{2, 0, 1}, calcRotateRight([]int{0, 1, 2}, 4))
}

func calcRotateRight(nums []int, k int) []int {
	head := createList(nums)
	head = rotateRight(head, k)
	res := printList(head)
	return res
}

func TestRotateRight2(t *testing.T) {
	assert.Equal(t, []int{4, 5, 1, 2, 3}, calcRotateRight2([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{2, 0, 1}, calcRotateRight2([]int{0, 1, 2}, 4))
}

func calcRotateRight2(nums []int, k int) []int {
	head := createList(nums)
	head = rotateRight2(head, k)
	res := printList(head)
	return res
}

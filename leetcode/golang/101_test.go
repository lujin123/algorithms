package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{1, 2, 2, 3, 4, 4, 3})))

	assert.Equal(t, false, isSymmetric(createFulllikeTreeWithArray([]int{1, 2, 2, 0, 3, 0, 3})))

	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{1})))
	assert.Equal(t, true, isSymmetric(createFulllikeTreeWithArray([]int{})))
}

func TestIsSymmetric2(t *testing.T) {
	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{1, 2, 2, 3, 4, 4, 3})))

	assert.Equal(t, false, isSymmetric2(createFulllikeTreeWithArray([]int{1, 2, 2, 0, 3, 0, 3})))

	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{1})))
	assert.Equal(t, true, isSymmetric2(createFulllikeTreeWithArray([]int{})))
}

func createFulllikeTreeWithArray(nums []int) (root *TreeNode) {
	n := len(nums)
	if n == 0 {
		return
	}
	var queue []*TreeNode

	index := 0

	queue = append(queue, &TreeNode{
		Val: nums[index],
	})

	index += 1

	for len(queue) > 0 && index < n {
		node := queue[0]
		queue = queue[1:]
		if root == nil {
			root = node
		}
		leftVal := nums[index]
		if leftVal != 0 {
			leftNode := &TreeNode{Val: leftVal}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		index += 1

		if index >= n {
			break
		}
		rightVal := nums[index]
		if rightVal != 0 {
			rightNode := &TreeNode{Val: rightVal}
			node.Right = rightNode
			queue = append(queue, rightNode)
		}
		index += 1
	}
	return
}

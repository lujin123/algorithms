package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	nodes := []int{1, 2, 9, 8, 3, 6}
	bst := &BinarySearchTree{}
	for i := range nodes {
		bst.Insert(nodes[i])
	}

	var res1 []int
	bst.InOrder(bst.root, &res1)
	assert.Equal(t, []int{1, 2, 3, 6, 8, 9}, res1)

	assert.Equal(t, 1, bst.Find(1).Elem)
	assert.Equal(t, 9, bst.FindMax(bst.root).Elem)
	assert.Equal(t, 1, bst.FindMin(bst.root).Elem)

	bst.Delete(3)
	var res2 []int
	bst.InOrder(bst.root, &res2)
	assert.Equal(t, []int{1, 2, 6, 8, 9}, res2)
}

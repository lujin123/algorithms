package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTree(t *testing.T) {
	avl := &AVLTree{}

	nodes := []int{1, 2, 9, 8, 3, 6}
	for i := range nodes {
		avl.Insert(nodes[i])
	}
	var result []int
	avl.PreOrder(avl.root, &result)
	assert.Equal(t, []int{3, 2, 1, 8, 6, 9}, result)

	avl.Delete(3)
	var res2 []int
	avl.PreOrder(avl.root, &res2)
	assert.Equal(t, []int{6, 2, 1, 8, 9}, res2)

	avl.Delete(1)
	var res3 []int
	avl.InOrder(avl.root, &res3)
	assert.Equal(t, []int{2, 6, 8, 9}, res3)

	avl.Delete(6)
	var res4 []int
	avl.InOrder(avl.root, &res4)
	assert.Equal(t, []int{2, 8, 9}, res4)

	avl.Delete(2)
	var res5 []int
	avl.InOrder(avl.root, &res5)
	assert.Equal(t, []int{8, 9}, res5)

	avl.Delete(8)
	var res6 []int
	avl.InOrder(avl.root, &res6)
	assert.Equal(t, []int{9}, res6)
}

func TestAVLTree2(t *testing.T) {
	avl := &AVLTree2{}

	nodes := []int{1, 2, 9, 8, 3, 6}
	for i := range nodes {
		avl.Add(nodes[i])
	}
	var result []int
	avl.InOrder(avl.root, &result)
	assert.Equal(t, []int{1, 2, 3, 6, 8, 9}, result)

	var res1 []int
	avl.Remove(3)
	avl.InOrder(avl.root, &res1)
	assert.Equal(t, []int{1, 2, 6, 8, 9}, res1)
}

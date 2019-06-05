package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRBTree(t *testing.T) {
	rb := NewRBTree()
	nodes := []int{10, 85, 15, 70, 20, 60, 30, 50, 65, 80, 90, 40, 5, 55}

	for _, x := range nodes {
		rb.Insert(x)
	}

	assert.Equal(t, true, rb.CheckRBTree())

	var res []int
	rb.InOrder(rb.root.right, &res)
	assert.Equal(t, []int{5, 10, 15, 20, 30, 40, 50, 55, 60, 65, 70, 80, 85, 90}, res)

	var res2 []int
	rb.Delete(30)
	assert.Equal(t, true, rb.CheckRBTree())
	rb.InOrder(rb.root.right, &res2)
	assert.Equal(t, []int{5, 10, 15, 20, 40, 50, 55, 60, 65, 70, 80, 85, 90}, res2)

	var res3 []int
	rb.Delete(20)
	assert.Equal(t, true, rb.CheckRBTree())
	rb.InOrder(rb.root.right, &res3)
	assert.Equal(t, []int{5, 10, 15, 40, 50, 55, 60, 65, 70, 80, 85, 90}, res3)

	rb.Delete(85)
	assert.Equal(t, true, rb.CheckRBTree())
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	matrix1 := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 5},
	}
	assert.Equal(t, true, searchMatrix(matrix1, 3))
	assert.Equal(t, false, searchMatrix(matrix1, 13))

	matrix2 := [][]int{
		[]int{},
	}

	assert.Equal(t, false, searchMatrix(matrix2, 1))
}

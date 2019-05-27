package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateImage(t *testing.T) {
	matrix1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	rotateImage(matrix1)
	assert.Equal(t, [][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}, matrix1)

	matrix2 := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotateImage(matrix2)
	assert.Equal(t, [][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}}, matrix2)

}

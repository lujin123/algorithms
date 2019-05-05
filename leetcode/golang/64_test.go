package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	assert.Equal(t, 7, minPathSum(grid))

	grid2 := [][]int{
		[]int{0},
	}
	assert.Equal(t, 0, minPathSum(grid2))
	grid3 := [][]int{
		[]int{1, 2, 5},
		[]int{3, 2, 1},
	}
	assert.Equal(t, 6, minPathSum(grid3))
}

func TestMinPathSum2(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	assert.Equal(t, 7, minPathSum2(grid))

	grid2 := [][]int{
		[]int{0},
	}
	assert.Equal(t, 0, minPathSum2(grid2))
	grid3 := [][]int{
		[]int{1, 2, 5},
		[]int{3, 2, 1},
	}
	assert.Equal(t, 6, minPathSum2(grid3))
}

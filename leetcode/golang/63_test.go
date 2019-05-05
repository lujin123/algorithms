package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 2, uniquePathsWithObstacles(grid))

	grid1 := [][]int{
		[]int{1, 0},
	}
	assert.Equal(t, 0, uniquePathsWithObstacles(grid1))

	grid2 := [][]int{
		[]int{0},
	}
	assert.Equal(t, 1, uniquePathsWithObstacles(grid2))
}

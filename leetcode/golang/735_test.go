package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	assert.Equal(t, []int{-2, -2, -2}, asteroidCollision([]int{-2, 1, -2, -2}))
	assert.Equal(t, []int{-2, -2, -2}, asteroidCollision([]int{-2, -2, 1, -2}))
	assert.Equal(t, []int{-2, -2, -2, -2}, asteroidCollision([]int{-2, -2, -2, -2}))
	assert.Equal(t, []int{5, 10}, asteroidCollision([]int{5, 10, -5}))
	assert.Equal(t, []int{}, asteroidCollision([]int{8, -8}))
	assert.Equal(t, []int{10}, asteroidCollision([]int{10, 2, -5}))
	assert.Equal(t, []int{-2, -1, 1, 2}, asteroidCollision([]int{-2, -1, 1, 2}))
}

func TestAsteroidCollision2(t *testing.T) {
	assert.Equal(t, []int{-2, -2, -2}, asteroidCollision2([]int{-2, 1, -2, -2}))
	assert.Equal(t, []int{-2, -2, -2}, asteroidCollision2([]int{-2, -2, 1, -2}))
	assert.Equal(t, []int{-2, -2, -2, -2}, asteroidCollision2([]int{-2, -2, -2, -2}))
	assert.Equal(t, []int{5, 10}, asteroidCollision2([]int{5, 10, -5}))
	assert.Equal(t, []int{}, asteroidCollision2([]int{8, -8}))
	assert.Equal(t, []int{10}, asteroidCollision2([]int{10, 2, -5}))
	assert.Equal(t, []int{-2, -1, 1, 2}, asteroidCollision2([]int{-2, -1, 1, 2}))
}

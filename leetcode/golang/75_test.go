package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	res := []int{0, 0, 1, 1, 2, 2}
	assert.Equal(t, res, nums)
}

func TestSortColors2(t *testing.T) {
	nums := []int{1, 2, 0}
	sortColors2(nums)
	assert.Equal(t, []int{0, 1, 2}, nums)
	nums1 := []int{2, 2}
	sortColors2(nums1)
	assert.Equal(t, []int{2, 2}, nums1)
	nums2 := []int{2, 0, 1}
	sortColors2(nums2)
	assert.Equal(t, []int{0, 1, 2}, nums2)
	nums3 := []int{2, 0, 2, 1, 1, 0}
	sortColors2(nums3)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums3)
}

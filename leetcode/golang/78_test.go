package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{1, 3}, []int{1, 2, 3}, []int{2, 3}, []int{3}, []int{}}, subsets([]int{1, 2, 3}))

	// 注意这个测试用例跑不通，就是二维数组内容被复用了，内容运行时被修改了
	assert.ElementsMatch(t, [][]int{[]int{9}, []int{9, 0}, []int{0}, []int{9, 3}, []int{9, 0, 3}, []int{0, 3}, []int{3}, []int{9, 5}, []int{9, 0, 5}, []int{0, 5}, []int{9, 3, 5}, []int{9, 0, 3, 5}, []int{0, 3, 5}, []int{3, 5}, []int{5}, []int{9, 7}, []int{9, 0, 7}, []int{0, 7}, []int{9, 3, 7}, []int{9, 0, 3, 7}, []int{0, 3, 7}, []int{3, 7}, []int{9, 5, 7}, []int{9, 0, 5, 7}, []int{0, 5, 7}, []int{9, 3, 5, 7}, []int{9, 0, 3, 5, 7}, []int{0, 3, 5, 7}, []int{3, 5, 7}, []int{5, 7}, []int{7}, []int{}}, subsets([]int{9, 0, 3, 5, 7}))
}

func TestSubsets2(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{3}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}}, subsets2([]int{1, 2, 3}))

	assert.ElementsMatch(t, [][]int{[]int{}, []int{9}, []int{0}, []int{9, 0}, []int{3}, []int{9, 3}, []int{0, 3}, []int{9, 0, 3}, []int{5}, []int{9, 5}, []int{0, 5}, []int{9, 0, 5}, []int{3, 5}, []int{9, 3, 5}, []int{0, 3, 5}, []int{9, 0, 3, 5}, []int{7}, []int{9, 7}, []int{0, 7}, []int{9, 0, 7}, []int{3, 7}, []int{9, 3, 7}, []int{0, 3, 7}, []int{9, 0, 3, 7}, []int{5, 7}, []int{9, 5, 7}, []int{0, 5, 7}, []int{9, 0, 5, 7}, []int{3, 5, 7}, []int{9, 3, 5, 7}, []int{0, 3, 5, 7}, []int{9, 0, 3, 5, 7}}, subsets2([]int{9, 0, 3, 5, 7}))
}

func TestSubsets3(t *testing.T) {
	assert.ElementsMatch(t, [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{3}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}}, subsets3([]int{1, 2, 3}))

	assert.ElementsMatch(t, [][]int{[]int{}, []int{9}, []int{0}, []int{9, 0}, []int{3}, []int{9, 3}, []int{0, 3}, []int{9, 0, 3}, []int{5}, []int{9, 5}, []int{0, 5}, []int{9, 0, 5}, []int{3, 5}, []int{9, 3, 5}, []int{0, 3, 5}, []int{9, 0, 3, 5}, []int{7}, []int{9, 7}, []int{0, 7}, []int{9, 0, 7}, []int{3, 7}, []int{9, 3, 7}, []int{0, 3, 7}, []int{9, 0, 3, 7}, []int{5, 7}, []int{9, 5, 7}, []int{0, 5, 7}, []int{9, 0, 5, 7}, []int{3, 5, 7}, []int{9, 3, 5, 7}, []int{0, 3, 5, 7}, []int{9, 0, 3, 5, 7}}, subsets3([]int{9, 0, 3, 5, 7}))
}

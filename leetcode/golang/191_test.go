package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(11))
	assert.Equal(t, 1, hammingWeight(256))
}

func TestHammingWeight2(t *testing.T) {
	assert.Equal(t, 3, hammingWeight2(11))
	assert.Equal(t, 1, hammingWeight2(256))
}

func TestHammingWeight3(t *testing.T) {
	assert.Equal(t, 3, hammingWeight3(11))
	assert.Equal(t, 1, hammingWeight3(256))
}

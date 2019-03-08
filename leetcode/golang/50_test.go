package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	assert.Equal(t, 1.00000, myPow(2.00000, 0))
	assert.Equal(t, 1024.00000, myPow(2.00000, 10))
	assert.Equal(t, 9.261000000000001, myPow(2.10000, 3))
	assert.Equal(t, 0.25000, myPow(2.00000, -2))
}

package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, -42, myAtoi("   -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, math.MinInt32, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, 0, myAtoi("+-2"))
	assert.Equal(t, math.MaxInt32, myAtoi("9223372036854775808"))
	assert.Equal(t, math.MinInt32, myAtoi("-9223372036854775809"))
	assert.Equal(t, math.MaxInt32, myAtoi("18446744073709551617"))
	assert.Equal(t, math.MinInt32, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("    0000000000000   "))
	assert.Equal(t, math.MinInt32, myAtoi("-6147483648"))
	assert.Equal(t, math.MaxInt32, myAtoi("20000000000000000000"))
	assert.Equal(t, 0, myAtoi("0-1"))
	assert.Equal(t, math.MaxInt32, myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	assert.Equal(t, 12345678, myAtoi("  0000000000012345678"))
	assert.Equal(t, 0, myAtoi("   "))
	assert.Equal(t, 1, myAtoi("+1"))
}

func TestMyAtoi2(t *testing.T) {
	assert.Equal(t, -42, myAtoi2("   -42"))
	assert.Equal(t, 4193, myAtoi2("4193 with words"))
	assert.Equal(t, math.MinInt32, myAtoi2("-91283472332"))
	assert.Equal(t, 0, myAtoi2("words and 987"))
	assert.Equal(t, 0, myAtoi2("+-2"))
	assert.Equal(t, math.MaxInt32, myAtoi2("9223372036854775808"))
	assert.Equal(t, math.MinInt32, myAtoi2("-9223372036854775809"))
	assert.Equal(t, math.MaxInt32, myAtoi2("18446744073709551617"))
	assert.Equal(t, math.MinInt32, myAtoi2("-91283472332"))
	assert.Equal(t, 0, myAtoi2("    0000000000000   "))
	assert.Equal(t, math.MinInt32, myAtoi2("-6147483648"))
	assert.Equal(t, math.MaxInt32, myAtoi2("20000000000000000000"))
	assert.Equal(t, 0, myAtoi2("0-1"))
	assert.Equal(t, math.MaxInt32, myAtoi2("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	assert.Equal(t, 12345678, myAtoi2("  0000000000012345678"))
	assert.Equal(t, 0, myAtoi2("   "))
	assert.Equal(t, 1, myAtoi2("+1"))
}

package leetcode

import (
	"strings"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false

// 思路：
// 1. 由于没有找到go中用于将byte转成小写的方法，所以先全部将字符串转成小写，否则直接在比较的时候转换即可
// 2. 可以先将不符合要求的字符去掉，在进行比较

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphaNum(s[i]) {
			i++
			continue
		}
		if !isAlphaNum(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

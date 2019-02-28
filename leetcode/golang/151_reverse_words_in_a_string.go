package leetcode

import "strings"

// doc: https://leetcode-cn.com/problems/reverse-words-in-a-string/
// Given an input string, reverse the string word by word.

// Example 1:

// Input: "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: "  hello world!  "
// Output: "world! hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:

// Input: "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

// Note:

// A word is defined as a sequence of non-space characters.
// Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
// You need to reduce multiple spaces between two words to a single space in the reversed string.

// Follow up:

// For C programmers, try to solve it in-place in O(1) extra space.

// 思路：
// 这题简单的直接依照空格分割开字符串即可
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var results []string
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		if val != "" {
			results = append(results, val)
		}
	}
	return strings.Join(results, " ")
}

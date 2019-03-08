package leetcode

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/

/**
找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。

示例 1:

输入:
s = "aaabb", k = 3

输出:
3

最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2:

输入:
s = "ababbc", k = 2

输出:
5

最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
*/

// 主要是找出小于k次的字符，然后将字符串进行分割，递归进行，每次分割都要进行判断是否符合要求字符串，符合则计算长度
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	a := make([]int, 26)
	for _, v := range s {
		a[v-'a'] += 1
	}

	max := 0
	var dp []rune
	for _, v := range s {
		if a[v-'a'] >= k {
			dp = append(dp, v)
		} else {
			max = calMax(dp, k, max)
			dp = dp[:0]
		}
	}
	max = calMax(dp, k, max)
	return max
}

func calMax(dp []rune, k, max int) int {
	n := len(dp)
	if n >= k {
		b := make([]int, 26)
		for _, i := range dp {
			b[i-'a']++
		}

		flag := true
		for _, i := range b {
			if i > 0 && i < k {
				flag = false
				break
			}
		}
		if flag {
			if n > max {
				max = n
			}
		} else {
			temp := longestSubstring(string(dp), k)
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

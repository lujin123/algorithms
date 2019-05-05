package leetcode

// doc: https://leetcode-cn.com/problems/permutation-in-string/
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

// 换句话说，第一个字符串的排列之一是第二个字符串的子串。

// 示例1:

// 输入: s1 = "ab" s2 = "eidbaooo"
// 输出: True
// 解释: s2 包含 s1 的排列之一 ("ba").

// 示例2:

// 输入: s1= "ab" s2 = "eidboaoo"
// 输出: False

// 注意：

// 输入的字符串只包含小写字母
// 两个字符串的长度都在 [1, 10,000] 之间

// 思路：这题直接暴力求解估计是要TTL，所以只能想点别的方法
// 1. 统计两个数组的词频，每次新增一个字母，就要判断下是否满足条件
func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}
	diff := make([]int, 26)
	for i := 0; i < n1; i++ {
		diff[s1[i]-'a']--
		diff[s2[i]-'a']++
	}
	for i := n1; i < n2; i++ {
		if isZero(diff) {
			return true
		}
		// 将多加的首字符减掉
		diff[s2[i-n1]-'a']--
		diff[s2[i]-'a']++
	}
	return isZero(diff)
}

func isZero(diff []int) bool {
	for _, v := range diff {
		if v != 0 {
			return false
		}
	}
	return true
}

// 2. 在大佬极致的优化下，成功的去掉了对数组的遍历判断，直接对两个整数判断是否相等即可
func checkInclusion2(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}

	h1 := 0
	h2 := 0
	for i := 0; i < n1; i++ {
		c1 := s1[i] - 'a'
		c2 := s2[i] - 'a'
		h1 += 1 << c1
		h2 += 1 << c2
	}

	if h1 == h2 {
		return true
	}

	for i := n1; i < n2; i++ {
		cb := s2[i-n1] - 'a'
		ce := s2[i] - 'a'
		h2 += (1 << ce) - (1 << cb)
		if h1 == h2 {
			return true
		}
	}

	return false
}

package leetcode

// https://leetcode-cn.com/problems/word-break/
/**
139. 单词拆分

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

// 1. 动态规划，遍历字符串，对于每个位置的元素表是否与前面的元素组成单词存在于字典中
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				var flag bool
				for k := range wordDict {
					if wordDict[k] == s[j:i] {
						flag = true
						break
					}
				}
				if flag {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}

// 2. 动态规划的优化版本
// 因为字典数组每次都要遍历，最少也O(logN)，所以变成字典用空间换取时间
// 并且每次也不用从0开始查询，只要从字段中最长的元素开始遍历即可
func wordBreak2(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	var maxLength int
	wordDictMap := make(map[string]bool)
	for i := range wordDict {
		n := len(wordDict[i])
		if n > maxLength {
			maxLength = n
		}
		wordDictMap[wordDict[i]] = true
	}

	for i := 1; i <= n; i++ {
		j := i - maxLength
		if j < 0 {
			j = 0
		}
		for ; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// 3. 利用前缀树，也可以实现，首先要将字典中的字符串变成一颗前缀树，然后遍历字符串在前缀树中查找对应的元素，
// 直到叶子节点位置，就算分割出一个单词，一直到结尾都找到了就算分割完成，否则在某个非叶子节点找不到元素，就表示字符串无法分割
// 可能空间复杂度比较高，但是时间复杂度还行应该
// 上面都是我猜的，我没实现...哈哈
func wordBreak3(s string, wordDict []string) bool {
	// todo
	return true
}

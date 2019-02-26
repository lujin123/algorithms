# Created by lujin at 4/3/2017
#
# 58. Length of Last Word
#
# Description:
#
# Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
# return the length of last word in the string.
#
# If the last word does not exist, return 0.
#
# Note: A word is defined as a character sequence consists of non-space characters only.
#
# For example,
# Given s = "Hello World",
# return 5.
#


class Solution(object):
    def lengthOfLastWord(self, s):
        """
        从字符串最后开始统计，并且一开始的空格需要去掉

        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        k = 0
        for i in range(len(s)-1, -1, -1):
            if s[i] != ' ':
                k += 1
            elif s[i] == ' ' and k:
                break
        return k

    def test(self):
        s = 'hello world  '
        # s = 'a '
        print(self.lengthOfLastWord(s))


if __name__ == '__main__':
    Solution().test()

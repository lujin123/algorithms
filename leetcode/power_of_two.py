# Created by lujin at 6/3/2017

# 231. Power of Two
#
# Description:
#
# Given an integer, write a function to determine if it is a power of two.
#


class Solution(object):
    def isPowerOfTwo(self, n):
        """
        由于符合要求的数n的二进制表示中只有一个1，而且是在最高位，
        所以 n-1 之后就是将所有的0位变成1，1变成0，之后相与，结果应该就是是 0 了

        :type n: int
        :rtype: bool
        """
        return n > 0 and not (n & (n - 1))

    def isPowerOfTwo2(self, n):
        """
        分析后可以发现，2的幂的数的二进制都是只有一个1，所以只需要统计下二进制表示有几个1了，
        然后只有一个的话就是 True

        :param n:
        :return:
        """
        count = 0
        while n > 0:
            count += (n & 1)
            n >>= 1
        return count == 1

    def test(self):
        print(self.isPowerOfTwo(2))
        print(self.isPowerOfTwo(1))
        print(self.isPowerOfTwo(0))
        print(self.isPowerOfTwo(-1))
        print(self.isPowerOfTwo(3))
        print(self.isPowerOfTwo(8))

        print(self.isPowerOfTwo2(7))
        print(self.isPowerOfTwo2(8))


if __name__ == '__main__':
    Solution().test()

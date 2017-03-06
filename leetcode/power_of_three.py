# Created by lujin at 6/3/2017
#
# 326. Power of Three
#
# Description:
#
# Given an integer, write a function to determine if it is a power of three.
#
# Follow up:
# Could you do it without using any loop / recursion?
#
import math


class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        return n > 0 and not (1162261467 % n)

    def isPowerOfThree2(self, n):
        """
        n = 3^x ==> x = log(n) / log(3)
        判断 x 是否为 0 即可

        :type n: int
        :rtype: bool
        """
        return n > 0 and not int(math.log10(n)-math.log10(3))

    def test(self):
        print(self.isPowerOfThree(3))
        print(self.isPowerOfThree(9))
        print(self.isPowerOfThree(0))
        print(self.isPowerOfThree(1))
        print('===========================')
        print(self.isPowerOfThree2(3))
        print(self.isPowerOfThree2(9))
        print(self.isPowerOfThree2(0))
        print(self.isPowerOfThree2(1))
        pass


if __name__ == '__main__':
    Solution().test()

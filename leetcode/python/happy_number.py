# Created by lujin at 6/3/2017
#
# 202. Happy Number
#
# Description:
#
# Write an algorithm to determine if a number is "happy".
#
# A happy number is a number defined by the following process: Starting with any positive integer,
# replace the number by the sum of the squares of its digits, and repeat the process until the number
# equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
# Those numbers for which this process ends in 1 are happy numbers.
#
# Example: 19 is a happy number
#
# 12 + 92 = 82
# 82 + 22 = 68
# 62 + 82 = 100
# 12 + 02 + 02 = 1
#
import math


class Solution(object):
    def isHappy(self, n):
        """
        找寻happy number，如果最后结果为1，则true，否则会进入循环，而且循环的数字是固定的那么几个...

        :type n: int
        :rtype: bool
        """
        if n <= 0:
            return False

        loop = []
        while n != 1:
            num = 0
            while n:
                num += int(math.pow(n % 10, 2))
                n //= 10
            n = num
            if n in loop:
                break
            else:
                loop.append(num)
        return n == 1

    def isHappy2(self, n):
        """
        重复的数字中必定有4，所以可以直接判断，含有4，必定不是happy number

        :param n:
        :return:
        """
        if n <= 0:
            return False

        while n != 1 and n != 4:
            num = 0
            while n:
                num += int(math.pow(n % 10, 2))
                n //= 10
            n = num
        return n == 1

    def test(self):
        print(self.isHappy(1))
        print(self.isHappy(7))
        print(self.isHappy(11))
        print(self.isHappy(100))


if __name__ == '__main__':
    Solution().test()

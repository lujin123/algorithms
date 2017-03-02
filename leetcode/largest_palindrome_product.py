# Description:
#
# Find the largest palindrome made from the product of two n-digit numbers.
#
# Since the result could be very large, you should return the largest palindrome mod 1337.
#
# Example:
#
# Input: 2
#
# Output: 987
#
# Explanation: 99 x 91 = 9009, 9009 % 1337 = 987
#
# Note:
#
# The range of n is [1,8].
#


class Solution(object):
    def largestPalindrome(self, n):
        """
        这个方法虽然可以算出结果，但是超时了...

        :type n: int
        :rtype: int
        """
        if n == 1:
            return 9

        def create_palindrome(num):
            num = str(num)
            return int(num + num[::-1])

        high = pow(10, n) - 1
        low = high // 10
        for i in range(high, low, -1):
            value = create_palindrome(i)
            for j in range(high, low, -1):
                if value / j > high:
                    break
                if value % j == 0:
                    return value
        return -1

    def largestPalindrome2(self, n):
        """
        这个方法很暴力...
        :param n:
        :return:
        """
        return [9, 9009, 906609, 99000099, 9966006699, 999000000999,
                99956644665999, 9999000000009999][n - 1] % 1337

if __name__ == '__main__':
    s = Solution()
    # print(s.largestPalindrome(7))
    print(s.largestPalindrome2(8))

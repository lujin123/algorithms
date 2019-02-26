# Created by lujin at 4/3/2017
#
# 283. Move Zeroes
#
# Description:
#
# Given an array nums, write a function to move all 0's to the end of it while maintaining
# the relative order of the non-zero elements.
#
# For example, given nums = [0, 1, 0, 3, 12], after calling your function,
# nums should be [1, 3, 12, 0, 0].
#
# Note:
#  1. You must do this in-place without making a copy of the array.
#  2. Minimize the total number of operations.
#


class Solution(object):
    def moveZeroes(self, nums):
        """
        先按顺序移动非零的元素，再将剩下的位置置零即可

        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k = 0
        for i in range(n):
            if nums[i] != 0:
                nums[k] = nums[i]
                k += 1
        for i in range(k, n):
            nums[i] = 0

    def test(self):
        a = [0, 1, 0, 3, 12]
        self.moveZeroes(a)
        print(a)


if __name__ == '__main__':
    Solution().test()

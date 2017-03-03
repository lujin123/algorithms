# Created by lujin at 3/3/2017
#
# 27. Remove Element
#
# Description:
#
# Given an array and a value, remove all instances of that value in place and return the new length.
#
# Do not allocate extra space for another array, you must do this in place with constant memory.
#
# The order of elements can be changed. It doesn't matter what you leave beyond the new length.
#
# Example:
# Given input array nums = [3,2,2,3], val = 3
#
# Your function should return length = 2, with the first two elements of nums being 2.
#


class Solution(object):
    def removeElement(self, nums, val):
        """
        这个题目很奇怪，基本可以啥也不干，直接遍历一遍数组，统计一下不相同值的个数就可以了嘛...

        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        k = 0
        for i, ele in enumerate(nums):
            if val != ele:
                nums[k] = nums[i]
                k += 1
        return k

    def test(self):
        a = [3, 2, 2, 3]
        print(self.removeElement(a, 3))


if __name__ == '__main__':
    Solution().test()

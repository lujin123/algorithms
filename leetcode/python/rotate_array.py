# Created by lujin at 2/3/2017
#
# Rotate Array
#
# Description:
#
# Rotate an array of n elements to the right by k steps.
#
# For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
#
# Note:
# Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
#


class Solution(object):
    def rotate(self, nums, k):
        """
        copy一个数组然后算好移动后的位置直接填充到对应的位置即可

        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        if not nums:
            return
        n = len(nums)
        cp = nums[::]
        for i, e in enumerate(cp):
            nums[(i + k) % n] = cp[i]

    def rotate2(self, nums, k):
        """
        时间复杂度为O(n),空间复杂度O(1)，算好位置后替换原来位置的值，然后计算完为止

        :param nums:
        :param k:
        :return:
        """
        if not nums:
            return
        n = len(nums)
        index = 0
        distance = 0
        ele = nums[index]
        for _ in range(n):
            index = (index + k) % n
            nums[index], ele = ele, nums[index]
            distance = (distance + k) % n
            if not distance:
                index = (index + 1) % n
                ele = nums[index]

    def test(self):
        nums = [1, 2, 3, 4, 5, 6, 7]
        # nums = [1, 2, 3, 4, 5, 6]
        # nums = [2147483647,-2147483648,33,219,0]
        #
        # self.rotate(nums, 2)
        self.rotate2(nums, 2)
        print(nums)


if __name__ == '__main__':
    s = Solution()
    s.test()

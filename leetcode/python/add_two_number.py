# Description
#
# You are given two non-empty linked lists representing two non-negative integers.
# The digits are stored in reverse order and each of their nodes contain a single digit.
# Add the two numbers and return it as a linked list.
#
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.
#
# Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
# Output: 7 -> 0 -> 8
#


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def add_two_numbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if not l1:
            return l2
        if not l2:
            return l1
        p = None
        tail = None
        p1 = l1
        p2 = l2
        flag = 0
        while p1 and p2:
            val = p1.val + p2.val + flag
            rest = val - 10
            if rest >= 0:
                flag = 1
                val = rest
            else:
                flag = 0

            node = ListNode(val)
            if p:
                tail.next = node
            else:
                p = node
            tail = node
            p1 = p1.next
            p2 = p2.next

        q = tmp = p1 if p1 else p2
        tail.next = q
        while flag and tmp:
            val = tmp.val + flag
            rest = val - 10
            if rest >= 0:
                tmp.val = rest
            else:
                tmp.val = val
                flag = 0
                break
            tail = tmp
            tmp = tmp.next
        if flag:
            tail.next = ListNode(1)
        return p

    def test(self):
        # v1 = [9, 2, 3]
        # v2 = [2, 3, 0, 9, 9, 1]
        v1 = [9, 8]
        v2 = [1]
        l1 = None
        l2 = None
        p = None
        for i in v1:
            node = ListNode(i)
            if l1:
                p.next = node
            else:
                l1 = node
            p = node
        p = None
        for i in v2:
            node = ListNode(i)
            if l2:
                p.next = node
            else:
                l2 = node
            p = node
        p = self.add_two_numbers(l1, l2)
        while p:
            print(str(p.val)+'->', end='')
            p = p.next

if __name__ == '__main__':
    s = Solution()
    s.test()
